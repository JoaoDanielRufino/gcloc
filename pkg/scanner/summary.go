package scanner

import (
	"sort"
	"strconv"
	"strings"
)

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

func (s *Summary) OrderByLanguage(sortOrder string) {
	sortedLanguages := s.sortLanguages(sortOrder)

	for _, language := range sortedLanguages {
		result := s.Languages[language]
		s.OrderedResultByLanguage = append(s.OrderedResultByLanguage, []string{
			language,
			strconv.Itoa(result.Lines),
			strconv.Itoa(result.BlankLines),
			strconv.Itoa(result.Comments),
			strconv.Itoa(result.CodeLines),
		})
	}
}

func (s *Summary) OrderByCodeLines(sortOrder string) {
	var aux [][]string

	for language, result := range s.Languages {
		aux = append(aux, []string{
			language,
			strconv.Itoa(result.Lines),
			strconv.Itoa(result.BlankLines),
			strconv.Itoa(result.Comments),
			strconv.Itoa(result.CodeLines),
		})
	}

	order := strings.ToUpper(sortOrder)
	if order == "ASC" {
		sort.Slice(aux, func(i, j int) bool {
			a, _ := strconv.Atoi(aux[i][4])
			b, _ := strconv.Atoi(aux[j][4])
			return a < b
		})
	} else if order == "DESC" {
		sort.Slice(aux, func(i, j int) bool {
			a, _ := strconv.Atoi(aux[i][4])
			b, _ := strconv.Atoi(aux[j][4])
			return a > b
		})
	}

	s.OrderedResultByLanguage = aux
}

func (s *Summary) sortLanguages(sortOrder string) []string {
	var languages []string

	for language := range s.Languages {
		languages = append(languages, language)
	}

	order := strings.ToUpper(sortOrder)
	if order == "ASC" {
		sort.Strings(languages)
	} else if order == "DESC" {
		sort.Sort(sort.Reverse(sort.StringSlice(languages)))
	}

	return languages
}
