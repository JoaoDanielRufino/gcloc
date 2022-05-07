package reporter

import "github.com/JoaoDanielRufino/gcloc/pkg/sorter"

type Reporter interface {
	GenerateReportByLanguage(summary *sorter.SortedSummary) error
	GenerateReportByFile(summary *sorter.SortedSummary) error
}
