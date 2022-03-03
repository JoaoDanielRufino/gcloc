package gcloc

import (
	"fmt"

	"github.com/JoaoDanielRufino/gcloc/pkg/analyzer"
	"github.com/JoaoDanielRufino/gcloc/pkg/filesystem"
	"github.com/JoaoDanielRufino/gcloc/pkg/scanner"
)

type Params struct {
	Path         string
	ExcludePaths []string
}

type LanguageInfo struct {
	LineComments     []string
	MultiLineComment [][]string
}

type GCloc struct {
	Params              Params
	SupportedExtensions map[string]string
	SupprotedLanguages  map[string]LanguageInfo
	analyzer            *analyzer.Analyzer
	scanner             *scanner.Scanner
}

func NewGCloc(params Params, extensions map[string]string, languages map[string]LanguageInfo) (*GCloc, error) {
	excludePaths, err := filesystem.GetExcludePaths(params.Path, params.ExcludePaths)
	if err != nil {
		return &GCloc{}, err
	}

	analyzer := analyzer.NewAnalyzer(
		params.Path,
		excludePaths,
		extensions,
	)

	scanner := scanner.NewScanner()

	return &GCloc{
		Params:              params,
		SupportedExtensions: extensions,
		SupprotedLanguages:  languages,
		analyzer:            analyzer,
		scanner:             scanner,
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

	fmt.Println(scanResult)

	return nil
}
