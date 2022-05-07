package file

import (
	"sort"
	"strings"

	"github.com/JoaoDanielRufino/gcloc/pkg/scanner"
	"github.com/JoaoDanielRufino/gcloc/pkg/sorter"
)

type FileSorter struct {
	SortOrder string
}

func (f FileSorter) OrderByLanguage(summary *scanner.Summary) *sorter.SortedSummary {
	results := []sorter.Result{}

	for _, result := range summary.Files {
		results = append(results, sorter.Result{
			Name:       result.Path,
			Lines:      result.Lines,
			CodeLines:  result.CodeLines,
			BlankLines: result.BlankLines,
			Comments:   result.Comments,
		})
	}

	order := strings.ToUpper(f.SortOrder)
	if order == "ASC" {
		sort.Slice(results, func(i, j int) bool {
			a := results[i].Name
			b := results[j].Name
			return a < b
		})
	} else if order == "DESC" {
		sort.Slice(results, func(i, j int) bool {
			a := results[i].Name
			b := results[j].Name
			return a > b
		})
	}

	return &sorter.SortedSummary{
		Results:         results,
		TotalLines:      summary.TotalLines,
		TotalCodeLines:  summary.TotalCodeLines,
		TotalBlankLines: summary.TotalBlankLines,
		TotalComments:   summary.TotalComments,
	}
}

func (f FileSorter) OrderByCodeLines(summary *scanner.Summary) *sorter.SortedSummary {
	results := []sorter.Result{}

	for _, result := range summary.Files {
		results = append(results, sorter.Result{
			Name:       result.Path,
			Lines:      result.Lines,
			CodeLines:  result.CodeLines,
			BlankLines: result.BlankLines,
			Comments:   result.Comments,
		})
	}

	order := strings.ToUpper(f.SortOrder)
	if order == "ASC" {
		sort.Slice(results, func(i, j int) bool {
			a := results[i].CodeLines
			b := results[j].CodeLines
			return a < b
		})
	} else if order == "DESC" {
		sort.Slice(results, func(i, j int) bool {
			a := results[i].CodeLines
			b := results[j].CodeLines
			return a > b
		})
	}

	return &sorter.SortedSummary{
		Results:         results,
		TotalLines:      summary.TotalLines,
		TotalCodeLines:  summary.TotalCodeLines,
		TotalBlankLines: summary.TotalBlankLines,
		TotalComments:   summary.TotalComments,
	}
}

func (f FileSorter) OrderByLines(summary *scanner.Summary) *sorter.SortedSummary {
	results := []sorter.Result{}

	for _, result := range summary.Files {
		results = append(results, sorter.Result{
			Name:       result.Path,
			Lines:      result.Lines,
			CodeLines:  result.CodeLines,
			BlankLines: result.BlankLines,
			Comments:   result.Comments,
		})
	}

	order := strings.ToUpper(f.SortOrder)
	if order == "ASC" {
		sort.Slice(results, func(i, j int) bool {
			a := results[i].Lines
			b := results[j].Lines
			return a < b
		})
	} else if order == "DESC" {
		sort.Slice(results, func(i, j int) bool {
			a := results[i].Lines
			b := results[j].Lines
			return a > b
		})
	}

	return &sorter.SortedSummary{
		Results:         results,
		TotalLines:      summary.TotalLines,
		TotalCodeLines:  summary.TotalCodeLines,
		TotalBlankLines: summary.TotalBlankLines,
		TotalComments:   summary.TotalComments,
	}
}

func (f FileSorter) OrderByComments(summary *scanner.Summary) *sorter.SortedSummary {
	results := []sorter.Result{}

	for _, result := range summary.Files {
		results = append(results, sorter.Result{
			Name:       result.Path,
			Lines:      result.Lines,
			CodeLines:  result.CodeLines,
			BlankLines: result.BlankLines,
			Comments:   result.Comments,
		})
	}

	order := strings.ToUpper(f.SortOrder)
	if order == "ASC" {
		sort.Slice(results, func(i, j int) bool {
			a := results[i].Comments
			b := results[j].Comments
			return a < b
		})
	} else if order == "DESC" {
		sort.Slice(results, func(i, j int) bool {
			a := results[i].Comments
			b := results[j].Comments
			return a > b
		})
	}

	return &sorter.SortedSummary{
		Results:         results,
		TotalLines:      summary.TotalLines,
		TotalCodeLines:  summary.TotalCodeLines,
		TotalBlankLines: summary.TotalBlankLines,
		TotalComments:   summary.TotalComments,
	}
}

func (f FileSorter) OrderByBlankLines(summary *scanner.Summary) *sorter.SortedSummary {
	results := []sorter.Result{}

	for _, result := range summary.Files {
		results = append(results, sorter.Result{
			Name:       result.Path,
			Lines:      result.Lines,
			CodeLines:  result.CodeLines,
			BlankLines: result.BlankLines,
			Comments:   result.Comments,
		})
	}

	order := strings.ToUpper(f.SortOrder)
	if order == "ASC" {
		sort.Slice(results, func(i, j int) bool {
			a := results[i].BlankLines
			b := results[j].BlankLines
			return a < b
		})
	} else if order == "DESC" {
		sort.Slice(results, func(i, j int) bool {
			a := results[i].BlankLines
			b := results[j].BlankLines
			return a > b
		})
	}

	return &sorter.SortedSummary{
		Results:         results,
		TotalLines:      summary.TotalLines,
		TotalCodeLines:  summary.TotalCodeLines,
		TotalBlankLines: summary.TotalBlankLines,
		TotalComments:   summary.TotalComments,
	}
}
