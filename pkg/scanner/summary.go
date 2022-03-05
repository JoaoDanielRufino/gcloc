package scanner

type summary struct {
	Results         []scanResult
	TotalLines      int
	TotalCodeLines  int
	TotalBlankLines int
	TotalComments   int
}

func (sc *Scanner) Summary(results []scanResult) summary {
	summary := summary{Results: results}

	for _, result := range results {
		summary.TotalLines += result.Lines
		summary.TotalCodeLines += result.CodeLines
		summary.TotalBlankLines += result.BlankLines
		summary.TotalComments += result.Comments
	}

	return summary
}
