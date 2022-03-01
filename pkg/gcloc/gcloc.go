package gcloc

import "fmt"

type Params struct {
	Path        string
	ExcludeDirs []string
}

type GCloc struct {
	Params Params
}

func NewGCloc(params Params) GCloc {
	return GCloc{
		params,
	}
}

func (gc GCloc) Run() error {
	fmt.Println(gc.Params)

	return nil
}
