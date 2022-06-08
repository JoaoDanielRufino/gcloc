package json

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/JoaoDanielRufino/gcloc/pkg/sorter"
	"github.com/stretchr/testify/require"
)

func TestGenerateReportByLanguage(t *testing.T) {
	tests := []struct {
		name         string
		jsonReporter JsonReporter
		byLanguage   bool
		summary      *sorter.SortedSummary
	}{
		{
			name: "Should generate json report by language",
			jsonReporter: JsonReporter{
				OutputName: "results",
				OutputPath: os.TempDir(),
			},
			byLanguage: true,
			summary: &sorter.SortedSummary{
				Results: []sorter.Result{
					{
						Name:       "CPP",
						Lines:      100,
						CodeLines:  80,
						BlankLines: 10,
						Comments:   2,
					},
					{
						Name:       "Golang",
						Lines:      267,
						CodeLines:  217,
						BlankLines: 24,
						Comments:   18,
					},
				},
				FilesByLanguage: map[string]int{
					"Golang": 2,
					"CPP":    1,
				},
				TotalFiles:      3,
				TotalLines:      367,
				TotalCodeLines:  297,
				TotalBlankLines: 34,
				TotalComments:   20,
			},
		},
		{
			name: "Should generate json report by file",
			jsonReporter: JsonReporter{
				OutputName: "results",
				OutputPath: os.TempDir(),
			},
			summary: &sorter.SortedSummary{
				Results: []sorter.Result{
					{
						Name:       "/tmp/main.cpp",
						Lines:      100,
						CodeLines:  80,
						BlankLines: 10,
						Comments:   2,
					},
					{
						Name:       "/tmp/main.go",
						Lines:      167,
						CodeLines:  137,
						BlankLines: 14,
						Comments:   16,
					},
					{
						Name:       "/tmp/test.go",
						Lines:      100,
						CodeLines:  80,
						BlankLines: 10,
						Comments:   2,
					},
				},
				TotalLines:      367,
				TotalCodeLines:  297,
				TotalBlankLines: 34,
				TotalComments:   20,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			if tt.byLanguage {
				err = tt.jsonReporter.GenerateReportByLanguage(tt.summary)
			} else {
				err = tt.jsonReporter.GenerateReportByFile(tt.summary)
			}
			require.NoError(t, err)
			require.FileExists(t, filepath.Join(os.TempDir(), "results.json"))
		})
	}
}
