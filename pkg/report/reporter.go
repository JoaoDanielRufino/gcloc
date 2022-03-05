package report

import "github.com/JoaoDanielRufino/gcloc/pkg/scanner"

type Reporter interface {
	GenerateReport(summary *scanner.Summary) error
}
