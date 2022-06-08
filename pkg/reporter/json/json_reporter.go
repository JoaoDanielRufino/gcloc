package json

import "github.com/JoaoDanielRufino/gcloc/pkg/sorter"

type JsonReporter struct {
	OutputName string
	OutputPath string
}

func (j JsonReporter) GenerateReportByFile(summary *sorter.SortedSummary) error {
	return nil
}

func (j JsonReporter) GenerateReportByLanguage(summary *sorter.SortedSummary) error {
	return nil
}
