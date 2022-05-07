package sorter

import "github.com/JoaoDanielRufino/gcloc/pkg/scanner"

type Result struct {
	Name       string
	Lines      int
	CodeLines  int
	BlankLines int
	Comments   int
}

type SortedSummary struct {
	Results         []Result
	TotalLines      int
	TotalCodeLines  int
	TotalBlankLines int
	TotalComments   int
}

type Sorter interface {
	OrderByLanguage(summary *scanner.Summary) *SortedSummary
	OrderByCodeLines(summary *scanner.Summary) *SortedSummary
	OrderByLines(summary *scanner.Summary) *SortedSummary
	OrderByComments(summary *scanner.Summary) *SortedSummary
	OrderByBlankLines(summary *scanner.Summary) *SortedSummary
}
