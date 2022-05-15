package prompt

import (
	"os"
	"strconv"

	"github.com/JoaoDanielRufino/gcloc/pkg/sorter"
	"github.com/olekukonko/tablewriter"
)

type PromptReporter struct {
}

func (p PromptReporter) GenerateReportByLanguage(summary *sorter.SortedSummary) error {
	tableHeader := []string{
		"Language",
		"Files",
		"Lines",
		"Blank lines",
		"Comments",
		"Code lines",
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(tableHeader)
	table.SetBorder(false)
	table.SetAutoFormatHeaders(false)

	for _, file := range summary.Results {
		table.Append([]string{
			file.Name,
			strconv.Itoa(summary.FilesByLanguage[file.Name]),
			strconv.Itoa(file.Lines),
			strconv.Itoa(file.BlankLines),
			strconv.Itoa(file.Comments),
			strconv.Itoa(file.CodeLines),
		})
	}

	table.SetFooter([]string{
		"Total",
		strconv.Itoa(summary.TotalFiles),
		strconv.Itoa(summary.TotalLines),
		strconv.Itoa(summary.TotalBlankLines),
		strconv.Itoa(summary.TotalComments),
		strconv.Itoa(summary.TotalCodeLines),
	})

	table.Render()

	return nil
}

func (p PromptReporter) GenerateReportByFile(summary *sorter.SortedSummary) error {
	tableHeader := []string{
		"Path",
		"Lines",
		"Blank lines",
		"Comments",
		"Code lines",
	}

	p.printTable(tableHeader, summary)

	return nil
}

func (p PromptReporter) printTable(header []string, summary *sorter.SortedSummary) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetBorder(false)
	table.SetAutoFormatHeaders(false)

	for _, file := range summary.Results {
		table.Append([]string{
			file.Name,
			strconv.Itoa(file.Lines),
			strconv.Itoa(file.BlankLines),
			strconv.Itoa(file.Comments),
			strconv.Itoa(file.CodeLines),
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
}
