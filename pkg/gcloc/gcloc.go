package gcloc

import (
	"github.com/JoaoDanielRufino/gcloc/pkg/analyzer"
	"github.com/JoaoDanielRufino/gcloc/pkg/filesystem"
	"github.com/JoaoDanielRufino/gcloc/pkg/gcloc/language"
	"github.com/JoaoDanielRufino/gcloc/pkg/reporter"
	"github.com/JoaoDanielRufino/gcloc/pkg/reporter/prompt"
	"github.com/JoaoDanielRufino/gcloc/pkg/scanner"
	"github.com/JoaoDanielRufino/gcloc/pkg/sorter"
)

type Params struct {
	Path              string
	ExcludePaths      []string
	ExcludeExtensions []string
	ByFile            bool
	OrderByLang       bool
	OrderByFile       bool
	OrderByCode       bool
	OrderByLine       bool
	OrderByBlank      bool
	OrderByComment    bool
	Order             string
}

type GCloc struct {
	params   Params
	analyzer *analyzer.Analyzer
	scanner  *scanner.Scanner
	sorter   sorter.Sorter
	reporter reporter.Reporter
}

func NewGCloc(params Params, languages language.Languages) (*GCloc, error) {
	excludePaths, err := filesystem.GetExcludePaths(params.Path, params.ExcludePaths)
	if err != nil {
		return &GCloc{}, err
	}

	extensions := getExtensionsMap(languages)

	analyzer := analyzer.NewAnalyzer(
		params.Path,
		excludePaths,
		getExcludeExtensionsMap(params.ExcludeExtensions),
		extensions,
	)

	scanner := scanner.NewScanner(languages)

	sorter := getSorter(params.ByFile, params.Order)

	return &GCloc{
		params:   params,
		analyzer: analyzer,
		scanner:  scanner,
		sorter:   sorter,
		reporter: prompt.PromptReporter{},
	}, nil
}

func (gc *GCloc) Run() error {
	files, err := gc.analyzer.MatchingFiles()
	if err != nil {
		return err
	}

	scanResult, err := gc.scanner.Scan(files)
	if err != nil {
		return err
	}

	summary := gc.scanner.Summary(scanResult)

	sortedSummary := gc.sortSummary(summary)

	if gc.params.ByFile {
		err = gc.reporter.GenerateReportByFile(sortedSummary)
	} else {
		err = gc.reporter.GenerateReportByLanguage(sortedSummary)
	}

	return err
}

func (gc *GCloc) sortSummary(summary *scanner.Summary) *sorter.SortedSummary {
	params := gc.params

	if params.OrderByCode {
		return gc.sorter.OrderByCodeLines(summary)
	}

	if params.OrderByLang {
		return gc.sorter.OrderByLanguage(summary)
	}

	if params.OrderByLine {
		return gc.sorter.OrderByLines(summary)
	}

	if params.OrderByComment {
		return gc.sorter.OrderByComments(summary)
	}

	if params.OrderByBlank {
		return gc.sorter.OrderByBlankLines(summary)
	}

	if params.OrderByFile {
		if languageSorter, ok := gc.sorter.(sorter.LanguageSorter); ok {
			return languageSorter.OrderByFiles(summary)
		}
	}

	return gc.sorter.OrderByCodeLines(summary)
}

func (gc *GCloc) ChangeLanguages(languages language.Languages) {
	extensions := getExtensionsMap(languages)
	gc.scanner.SupportedLanguages = languages
	gc.analyzer.SupportedExtensions = extensions
}

func getExtensionsMap(languages language.Languages) map[string]string {
	extensions := map[string]string{}

	for language, languageInfo := range languages {
		for _, extension := range languageInfo.Extensions {
			extensions[extension] = language
		}
	}

	return extensions
}

func getExcludeExtensionsMap(excludeExtensionsParam []string) map[string]bool {
	excludeExtensions := map[string]bool{}

	for _, ex := range excludeExtensionsParam {
		excludeExtensions[ex] = true
	}

	return excludeExtensions
}

func getSorter(byFile bool, order string) sorter.Sorter {
	if byFile {
		return sorter.NewFileSorter(order)
	}

	return sorter.NewLanguageSorter(order)
}
