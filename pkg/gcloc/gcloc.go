package gcloc

import (
	"github.com/JoaoDanielRufino/gcloc/pkg/analyzer"
	"github.com/JoaoDanielRufino/gcloc/pkg/filesystem"
	"github.com/JoaoDanielRufino/gcloc/pkg/gcloc/language"
	"github.com/JoaoDanielRufino/gcloc/pkg/reporter"
	"github.com/JoaoDanielRufino/gcloc/pkg/reporter/prompt"
	"github.com/JoaoDanielRufino/gcloc/pkg/scanner"
	"github.com/JoaoDanielRufino/gcloc/pkg/sorter"
	languageSorter "github.com/JoaoDanielRufino/gcloc/pkg/sorter/language"
)

type Params struct {
	Path              string
	ExcludePaths      []string
	ExcludeExtensions []string
	ByFile            bool
	OrderByLang       bool
	OrderByCode       bool
	OrderByLine       bool
	OrderByBlank      bool
	OrderByComment    bool
	Order             string
}

type GCloc struct {
	params              Params
	supportedExtensions map[string]string
	supprotedLanguages  language.Languages
	analyzer            *analyzer.Analyzer
	scanner             *scanner.Scanner
	sorter              sorter.Sorter
	reporter            reporter.Reporter
}

func NewGCloc(params Params, extensions map[string]string, languages language.Languages) (*GCloc, error) {
	excludePaths, err := filesystem.GetExcludePaths(params.Path, params.ExcludePaths)
	if err != nil {
		return &GCloc{}, err
	}

	analyzer := analyzer.NewAnalyzer(
		params.Path,
		excludePaths,
		getExcludeExtensionsMap(params.ExcludeExtensions),
		extensions,
	)

	scanner := scanner.NewScanner(languages)

	sorter := languageSorter.LanguageSorter{
		SortOrder: params.Order,
	}

	return &GCloc{
		params:              params,
		supportedExtensions: extensions,
		supprotedLanguages:  languages,
		analyzer:            analyzer,
		scanner:             scanner,
		sorter:              sorter,
		reporter:            prompt.PromptReporter{},
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

	return gc.sorter.OrderByCodeLines(summary)
}

func (gc *GCloc) ChangeExtensions(extensions map[string]string) {
	gc.supportedExtensions = extensions
	gc.analyzer.SupportedExtensions = extensions
}

func (gc *GCloc) ChangeLanguages(languages language.Languages) {
	gc.supprotedLanguages = languages
	gc.scanner.SupportedLanguages = languages
}

func getExcludeExtensionsMap(excludeExtensionsParam []string) map[string]bool {
	excludeExtensions := map[string]bool{}

	for _, ex := range excludeExtensionsParam {
		excludeExtensions[ex] = true
	}

	return excludeExtensions
}
