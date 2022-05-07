package language

import (
	"sort"
	"strings"

	"github.com/JoaoDanielRufino/gcloc/pkg/scanner"
	"github.com/JoaoDanielRufino/gcloc/pkg/sorter"
)

type LanguageSorter struct {
	SortOrder string
}

func (l LanguageSorter) OrderByLanguage(summary *scanner.Summary) *sorter.SortedSummary {
	results := []sorter.Result{}
	sortedLanguages := l.sortLanguages(summary)

	for _, language := range sortedLanguages {
		result := summary.Languages[language]
		results = append(results, sorter.Result{
			Name:       language,
			Lines:      result.Lines,
			CodeLines:  result.CodeLines,
			BlankLines: result.BlankLines,
			Comments:   result.Comments,
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

func (l LanguageSorter) OrderByCodeLines(summary *scanner.Summary) *sorter.SortedSummary {
	results := []sorter.Result{}

	for language, result := range summary.Languages {
		results = append(results, sorter.Result{
			Name:       language,
			Lines:      result.Lines,
			CodeLines:  result.CodeLines,
			BlankLines: result.BlankLines,
			Comments:   result.Comments,
		})
	}

	order := strings.ToUpper(l.SortOrder)
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

func (l LanguageSorter) OrderByLines(summary *scanner.Summary) *sorter.SortedSummary {
	results := []sorter.Result{}

	for language, result := range summary.Languages {
		results = append(results, sorter.Result{
			Name:       language,
			Lines:      result.Lines,
			CodeLines:  result.CodeLines,
			BlankLines: result.BlankLines,
			Comments:   result.Comments,
		})
	}

	order := strings.ToUpper(l.SortOrder)
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

func (l LanguageSorter) OrderByComments(summary *scanner.Summary) *sorter.SortedSummary {
	results := []sorter.Result{}

	for language, result := range summary.Languages {
		results = append(results, sorter.Result{
			Name:       language,
			Lines:      result.Lines,
			CodeLines:  result.CodeLines,
			BlankLines: result.BlankLines,
			Comments:   result.Comments,
		})
	}

	order := strings.ToUpper(l.SortOrder)
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

func (l LanguageSorter) OrderByBlankLines(summary *scanner.Summary) *sorter.SortedSummary {
	results := []sorter.Result{}

	for language, result := range summary.Languages {
		results = append(results, sorter.Result{
			Name:       language,
			Lines:      result.Lines,
			CodeLines:  result.CodeLines,
			BlankLines: result.BlankLines,
			Comments:   result.Comments,
		})
	}

	order := strings.ToUpper(l.SortOrder)
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

func (l LanguageSorter) sortLanguages(summary *scanner.Summary) []string {
	var languages []string

	for language := range summary.Languages {
		languages = append(languages, language)
	}

	order := strings.ToUpper(l.SortOrder)
	if order == "ASC" {
		sort.Strings(languages)
	} else if order == "DESC" {
		sort.Sort(sort.Reverse(sort.StringSlice(languages)))
	}

	return languages
}
