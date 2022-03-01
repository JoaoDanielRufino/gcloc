package gcloc

import (
	"fmt"

	"github.com/JoaoDanielRufino/gcloc/internal/constants"
	"github.com/JoaoDanielRufino/gcloc/pkg/analyzer"
)

type Params struct {
	Path        string
	ExcludeDirs []string
}

type GCloc struct {
	Params       Params
	fileAnalyzer analyzer.Analyzer
}

func NewGCloc(params Params) GCloc {
	fileAnalyzer := analyzer.NewAnalyzer(
		params.Path,
		params.ExcludeDirs,
		constants.Extensions,
	)

	return GCloc{
		params,
		fileAnalyzer,
	}
}

func (gc GCloc) Run() error {
	files, err := gc.fileAnalyzer.Analyze()
	if err != nil {
		return err
	}

	fmt.Println(files)

	return nil
}
