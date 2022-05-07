package scanner

type languageResult struct {
	Lines      int
	CodeLines  int
	BlankLines int
	Comments   int
}

type fileResult struct {
	Path       string
	Lines      int
	CodeLines  int
	BlankLines int
	Comments   int
}

type Summary struct {
	Languages               map[string]*languageResult
	Files                   []fileResult
	OrderedResultByLanguage [][]string
	OrderedResultByFile     [][]string
	TotalLines              int
	TotalCodeLines          int
	TotalBlankLines         int
	TotalComments           int
}

func (sc *Scanner) Summary(results []scanResult) *Summary {
	summary := &Summary{Languages: make(map[string]*languageResult)}

	for _, result := range results {
		if value, ok := summary.Languages[result.Metadata.Language]; ok {
			value.Lines += result.Lines
			value.CodeLines += result.CodeLines
			value.BlankLines += result.BlankLines
			value.Comments += result.Comments
		} else {
			summary.Languages[result.Metadata.Language] = &languageResult{
				Lines:      result.Lines,
				CodeLines:  result.CodeLines,
				BlankLines: result.BlankLines,
				Comments:   result.Comments,
			}
		}

		summary.Files = append(summary.Files, fileResult{
			Path:       result.Metadata.FilePath,
			Lines:      result.Lines,
			CodeLines:  result.CodeLines,
			BlankLines: result.BlankLines,
			Comments:   result.Comments,
		})
		summary.TotalLines += result.Lines
		summary.TotalCodeLines += result.CodeLines
		summary.TotalBlankLines += result.BlankLines
		summary.TotalComments += result.Comments
	}

	return summary
}
