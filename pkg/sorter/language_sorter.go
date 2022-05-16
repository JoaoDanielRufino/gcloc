package sorter

import (
	"sort"
	"strings"

	"github.com/JoaoDanielRufino/gcloc/pkg/scanner"
)

type LanguageSorter struct {
	baseSorter
}

func NewLanguageSorter(sortOrder string) LanguageSorter {
	return LanguageSorter{
		baseSorter{
			strings.ToUpper(sortOrder),
		},
	}
}

func (l LanguageSorter) OrderByLanguage(summary *scanner.Summary) *SortedSummary {
	results := []Result{}
	sortedLanguages := l.sortLanguages(summary)

	for _, language := range sortedLanguages {
		result := summary.Languages[language]
		results = append(results, Result{
			Name:       language,
			Lines:      result.Lines,
			CodeLines:  result.CodeLines,
			BlankLines: result.BlankLines,
			Comments:   result.Comments,
		})
	}

	return &SortedSummary{
		Results:         results,
		FilesByLanguage: summary.FilesByLanguage,
		TotalFiles:      summary.TotalFiles,
		TotalLines:      summary.TotalLines,
		TotalCodeLines:  summary.TotalCodeLines,
		TotalBlankLines: summary.TotalBlankLines,
		TotalComments:   summary.TotalComments,
	}
}

func (l LanguageSorter) OrderByCodeLines(summary *scanner.Summary) *SortedSummary {
	results := l.getResults(summary)

	l.sortByCodeLines(results)

	return &SortedSummary{
		Results:         results,
		FilesByLanguage: summary.FilesByLanguage,
		TotalFiles:      summary.TotalFiles,
		TotalLines:      summary.TotalLines,
		TotalCodeLines:  summary.TotalCodeLines,
		TotalBlankLines: summary.TotalBlankLines,
		TotalComments:   summary.TotalComments,
	}
}

func (l LanguageSorter) OrderByLines(summary *scanner.Summary) *SortedSummary {
	results := l.getResults(summary)

	l.sortByLines(results)

	return &SortedSummary{
		Results:         results,
		FilesByLanguage: summary.FilesByLanguage,
		TotalFiles:      summary.TotalFiles,
		TotalLines:      summary.TotalLines,
		TotalCodeLines:  summary.TotalCodeLines,
		TotalBlankLines: summary.TotalBlankLines,
		TotalComments:   summary.TotalComments,
	}
}

func (l LanguageSorter) OrderByComments(summary *scanner.Summary) *SortedSummary {
	results := l.getResults(summary)

	l.sortByComments(results)

	return &SortedSummary{
		Results:         results,
		FilesByLanguage: summary.FilesByLanguage,
		TotalFiles:      summary.TotalFiles,
		TotalLines:      summary.TotalLines,
		TotalCodeLines:  summary.TotalCodeLines,
		TotalBlankLines: summary.TotalBlankLines,
		TotalComments:   summary.TotalComments,
	}
}

func (l LanguageSorter) OrderByBlankLines(summary *scanner.Summary) *SortedSummary {
	results := l.getResults(summary)

	l.sortByBlankLines(results)

	return &SortedSummary{
		Results:         results,
		FilesByLanguage: summary.FilesByLanguage,
		TotalFiles:      summary.TotalFiles,
		TotalLines:      summary.TotalLines,
		TotalCodeLines:  summary.TotalCodeLines,
		TotalBlankLines: summary.TotalBlankLines,
		TotalComments:   summary.TotalComments,
	}
}

func (l LanguageSorter) OrderByFiles(summary *scanner.Summary) *SortedSummary {
	results := l.getResults(summary)

	if l.sortOrder == "ASC" {
		sort.Slice(results, func(i, j int) bool {
			a := results[i].Name
			b := results[j].Name
			return summary.FilesByLanguage[a] < summary.FilesByLanguage[b]
		})
	} else if l.sortOrder == "DESC" {
		sort.Slice(results, func(i, j int) bool {
			a := results[i].Name
			b := results[j].Name
			return summary.FilesByLanguage[a] > summary.FilesByLanguage[b]
		})
	}

	return &SortedSummary{
		Results:         results,
		FilesByLanguage: summary.FilesByLanguage,
		TotalFiles:      summary.TotalFiles,
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

	if l.sortOrder == "ASC" {
		sort.Strings(languages)
	} else if l.sortOrder == "DESC" {
		sort.Sort(sort.Reverse(sort.StringSlice(languages)))
	}

	return languages
}

func (l LanguageSorter) getResults(summary *scanner.Summary) []Result {
	results := []Result{}

	for language, result := range summary.Languages {
		results = append(results, Result{
			Name:       language,
			Lines:      result.Lines,
			CodeLines:  result.CodeLines,
			BlankLines: result.BlankLines,
			Comments:   result.Comments,
		})
	}

	return results
}
