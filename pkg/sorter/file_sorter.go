package sorter

import (
	"sort"
	"strings"

	"github.com/JoaoDanielRufino/gcloc/pkg/scanner"
)

type FileSorter struct {
	baseSorter
}

func NewFileSorter(sortOrder string) FileSorter {
	return FileSorter{
		baseSorter{
			strings.ToUpper(sortOrder),
		},
	}
}

func (f FileSorter) OrderByLanguage(summary *scanner.Summary) *SortedSummary {
	results := f.getResults(summary)

	f.sortByFileName(results)

	return &SortedSummary{
		Results:         results,
		TotalLines:      summary.TotalLines,
		TotalCodeLines:  summary.TotalCodeLines,
		TotalBlankLines: summary.TotalBlankLines,
		TotalComments:   summary.TotalComments,
	}
}

func (f FileSorter) OrderByCodeLines(summary *scanner.Summary) *SortedSummary {
	results := f.getResults(summary)

	f.sortByCodeLines(results)

	return &SortedSummary{
		Results:         results,
		TotalLines:      summary.TotalLines,
		TotalCodeLines:  summary.TotalCodeLines,
		TotalBlankLines: summary.TotalBlankLines,
		TotalComments:   summary.TotalComments,
	}
}

func (f FileSorter) OrderByLines(summary *scanner.Summary) *SortedSummary {
	results := f.getResults(summary)

	f.sortByLines(results)

	return &SortedSummary{
		Results:         results,
		TotalLines:      summary.TotalLines,
		TotalCodeLines:  summary.TotalCodeLines,
		TotalBlankLines: summary.TotalBlankLines,
		TotalComments:   summary.TotalComments,
	}
}

func (f FileSorter) OrderByComments(summary *scanner.Summary) *SortedSummary {
	results := f.getResults(summary)

	f.sortByComments(results)

	return &SortedSummary{
		Results:         results,
		TotalLines:      summary.TotalLines,
		TotalCodeLines:  summary.TotalCodeLines,
		TotalBlankLines: summary.TotalBlankLines,
		TotalComments:   summary.TotalComments,
	}
}

func (f FileSorter) OrderByBlankLines(summary *scanner.Summary) *SortedSummary {
	results := f.getResults(summary)

	f.sortByBlankLines(results)

	return &SortedSummary{
		Results:         results,
		TotalLines:      summary.TotalLines,
		TotalCodeLines:  summary.TotalCodeLines,
		TotalBlankLines: summary.TotalBlankLines,
		TotalComments:   summary.TotalComments,
	}
}

func (f FileSorter) getResults(summary *scanner.Summary) []Result {
	results := []Result{}

	for _, result := range summary.Files {
		results = append(results, Result{
			Name:       result.Path,
			Lines:      result.Lines,
			CodeLines:  result.CodeLines,
			BlankLines: result.BlankLines,
			Comments:   result.Comments,
		})
	}

	return results
}

func (f FileSorter) sortByFileName(results []Result) {
	if f.sortOrder == "ASC" {
		sort.Slice(results, func(i, j int) bool {
			a := results[i].Name
			b := results[j].Name
			return a < b
		})
	} else if f.sortOrder == "DESC" {
		sort.Slice(results, func(i, j int) bool {
			a := results[i].Name
			b := results[j].Name
			return a > b
		})
	}
}
