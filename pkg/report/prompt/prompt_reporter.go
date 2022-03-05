package prompt

import (
	"os"
	"strconv"

	"github.com/JoaoDanielRufino/gcloc/pkg/scanner"
	"github.com/olekukonko/tablewriter"
)

type PromptReporter struct {
}

func (p PromptReporter) GenerateReport(summary *scanner.Summary) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"File path", "Lines", "Blank lines", "Comments", "Code lines"})
	table.SetBorder(false)
	table.SetAutoFormatHeaders(false)

	for _, result := range summary.Results {
		table.Append([]string{
			result.Metadata.FilePath,
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
