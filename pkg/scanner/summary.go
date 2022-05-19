package scanner

type LanguageResult struct {
	Lines      int
	CodeLines  int
	BlankLines int
	Comments   int
}

type FileResult struct {
	Path       string
	Lines      int
	CodeLines  int
	BlankLines int
	Comments   int
}

type Summary struct {
	Languages       map[string]*LanguageResult
	Files           []FileResult
	FilesByLanguage map[string]int
	TotalFiles      int
	TotalLines      int
	TotalCodeLines  int
	TotalBlankLines int
	TotalComments   int
}

func (sc *Scanner) Summary(results []scanResult) *Summary {
	summary := &Summary{
		Languages:       make(map[string]*LanguageResult),
		FilesByLanguage: make(map[string]int),
		TotalFiles:      len(results),
	}

	for _, result := range results {
		language := result.Metadata.Language
		if value, ok := summary.Languages[language]; ok {
			value.Lines += result.Lines
			value.CodeLines += result.CodeLines
			value.BlankLines += result.BlankLines
			value.Comments += result.Comments
		} else {
			summary.Languages[language] = &LanguageResult{
				Lines:      result.Lines,
				CodeLines:  result.CodeLines,
				BlankLines: result.BlankLines,
				Comments:   result.Comments,
			}
		}

		summary.Files = append(summary.Files, FileResult{
			Path:       result.Metadata.FilePath,
			Lines:      result.Lines,
			CodeLines:  result.CodeLines,
			BlankLines: result.BlankLines,
			Comments:   result.Comments,
		})
		summary.FilesByLanguage[language]++
		summary.TotalLines += result.Lines
		summary.TotalCodeLines += result.CodeLines
		summary.TotalBlankLines += result.BlankLines
		summary.TotalComments += result.Comments
	}

	return summary
}
