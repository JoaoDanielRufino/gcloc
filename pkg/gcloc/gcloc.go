package gcloc

import (
	"fmt"

	"github.com/JoaoDanielRufino/gcloc/internal/constants"
	"github.com/JoaoDanielRufino/gcloc/pkg/analyzer"
	"github.com/JoaoDanielRufino/gcloc/pkg/filesystem"
)

type Params struct {
	Path         string
	ExcludePaths []string
}

type GCloc struct {
	Params       Params
	fileAnalyzer *analyzer.Analyzer
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

	return GCloc{
		params,
		fileAnalyzer,
	}, nil
}

func (gc GCloc) Run() error {
	files, err := gc.fileAnalyzer.Analyze()
	if err != nil {
		return err
	}

	fmt.Println(files)

	return nil
}
