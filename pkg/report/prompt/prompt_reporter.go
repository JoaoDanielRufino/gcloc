package prompt

import (
	"os"
	"strconv"

	"github.com/JoaoDanielRufino/gcloc/pkg/scanner"
	"github.com/olekukonko/tablewriter"
)

type PromptReporter struct {
}

func (p PromptReporter) GenerateReportByLanguage(summary *scanner.Summary) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Language", "Lines", "Blank lines", "Comments", "Code lines"})
	table.SetBorder(false)
	table.SetAutoFormatHeaders(false)

	for language, result := range summary.Languages {
		table.Append([]string{
			language,
			strconv.Itoa(result.Lines),
			strconv.Itoa(result.BlankLines),
			strconv.Itoa(result.Comments),
			strconv.Itoa(result.CodeLines),
		})
	}

	table.SetFooter([]string{
		"Total",
		strconv.Itoa(summary.TotalLines),
		strconv.Itoa(summary.TotalBlankLines),
		strconv.Itoa(summary.TotalComments),
		strconv.Itoa(summary.TotalCodeLines),
	})

	table.Render()

	return nil
}

func (p PromptReporter) GenerateReportByFile(summary *scanner.Summary) error {
	return nil
}
