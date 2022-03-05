package report

import "github.com/JoaoDanielRufino/gcloc/pkg/scanner"

type Reporter interface {
	GenerateReportByLanguage(summary *scanner.Summary) error
	GenerateReportByFile(summary *scanner.Summary) error
}
