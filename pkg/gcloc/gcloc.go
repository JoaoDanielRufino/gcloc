package gcloc

import (
	"fmt"

	"github.com/JoaoDanielRufino/gcloc/internal/constants"
	"github.com/JoaoDanielRufino/gcloc/pkg/analyzer"
	"github.com/JoaoDanielRufino/gcloc/pkg/filesystem"
)

type Params struct {
	Path        string
	ExcludeDirs []string
}

type GCloc struct {
	Params       Params
	fileAnalyzer *analyzer.Analyzer
}

func NewGCloc(params Params) (GCloc, error) {
	excludeDirs, err := filesystem.GetExcludeDirs(params.Path, params.ExcludeDirs)
	if err != nil {
		return GCloc{}, err
	}

	fileAnalyzer := analyzer.NewAnalyzer(
		params.Path,
		excludeDirs,
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
