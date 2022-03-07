package subcommands

import (
	"os"
	"sort"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/JoaoDanielRufino/gcloc/internal/constants"
	"github.com/JoaoDanielRufino/gcloc/pkg/gcloc/language"
)

type commandRunnerFunc func(cmd *cobra.Command, args []string)

func NewLanguagesCmd() *cobra.Command {
	languagesCmd := &cobra.Command{
		Use:   "languages",
		Short: "Show gcloc supported languages",
		Run:   run(),
	}

	return languagesCmd
}

func run() commandRunnerFunc {
	return func(cmd *cobra.Command, args []string) {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Language", "Single Comments", "Multi Line Comments"})
		table.SetBorder(false)
		table.SetAutoFormatHeaders(false)

		sortedKeys := getSortedMapKeys(constants.Languages)

		for _, key := range sortedKeys {
			info := constants.Languages[key]
			table.Append([]string{
				key,
				formatSingleComment(info.LineComments),
				formatMultiLineComment(info.MultiLineComments),
			})
		}

		table.Render()
	}
}

func getSortedMapKeys(languages language.Languages) []string {
	var keys []string

	for k := range languages {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}

func formatSingleComment(singleComments []string) string {
	var comments string

	for _, comment := range singleComments {
		comments += " " + comment
	}

	return comments
}

func formatMultiLineComment(multiLineComments [][]string) string {
	var comments string

	for _, comment := range multiLineComments {
		comments += " " + comment[0] + " " + comment[1]
	}

	return comments
}
