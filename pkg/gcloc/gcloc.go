package gcloc

import (
	"github.com/JoaoDanielRufino/gcloc/pkg/analyzer"
	"github.com/JoaoDanielRufino/gcloc/pkg/filesystem"
	"github.com/JoaoDanielRufino/gcloc/pkg/gcloc/language"
	"github.com/JoaoDanielRufino/gcloc/pkg/reporter"
	"github.com/JoaoDanielRufino/gcloc/pkg/reporter/prompt"
	"github.com/JoaoDanielRufino/gcloc/pkg/scanner"
)

type Params struct {
	Path              string
	ExcludePaths      []string
	ExcludeExtensions []string
	ByFile            bool
}

type GCloc struct {
	params              Params
	supportedExtensions map[string]string
	supprotedLanguages  language.Languages
	analyzer            *analyzer.Analyzer
	scanner             *scanner.Scanner
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

	return &GCloc{
		params:              params,
		supportedExtensions: extensions,
		supprotedLanguages:  languages,
		analyzer:            analyzer,
		scanner:             scanner,
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

	if gc.params.ByFile {
		err = gc.reporter.GenerateReportByFile(summary)
	} else {
		err = gc.reporter.GenerateReportByLanguage(summary)
	}

	return err
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
	excludeExtensions := make(map[string]bool)

	for _, ex := range excludeExtensionsParam {
		excludeExtensions[ex] = true
	}

	return excludeExtensions
}
