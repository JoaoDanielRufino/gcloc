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
	Path         string
	ExcludePaths []string
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

	gc.reporter.GenerateReportByLanguage(summary)

	return nil
}

func (gc *GCloc) ChangeExtensions(extensions map[string]string) {
	gc.supportedExtensions = extensions
	gc.analyzer.SupportedExtensions = extensions
}

func (gc *GCloc) ChangeLanguages(languages language.Languages) {
	gc.supprotedLanguages = languages
	gc.scanner.SupportedLanguages = languages
}
