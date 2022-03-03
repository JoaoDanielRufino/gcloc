package gcloc

import (
	"fmt"

	"github.com/JoaoDanielRufino/gcloc/internal/constants"
	"github.com/JoaoDanielRufino/gcloc/pkg/analyzer"
	"github.com/JoaoDanielRufino/gcloc/pkg/filesystem"
	"github.com/JoaoDanielRufino/gcloc/pkg/scanner"
)

type Params struct {
	Path         string
	ExcludePaths []string
}

type GCloc struct {
	Params   Params
	analyzer *analyzer.Analyzer
	scanner  *scanner.Scanner
}

func NewGCloc(params Params) (GCloc, error) {
	excludePaths, err := filesystem.GetExcludePaths(params.Path, params.ExcludePaths)
	if err != nil {
		return GCloc{}, err
	}

	fileAnalyzer := analyzer.NewAnalyzer(
		params.Path,
		excludePaths,
		constants.Extensions,
	)

	scanner := scanner.NewScanner()

	return GCloc{
		params,
		fileAnalyzer,
		scanner,
	}, nil
}

func (gc GCloc) Run() error {
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
