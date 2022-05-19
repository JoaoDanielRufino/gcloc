package scanner

import (
	"testing"

	"github.com/JoaoDanielRufino/gcloc/internal/constants"
	"github.com/JoaoDanielRufino/gcloc/pkg/analyzer"
	"github.com/stretchr/testify/require"
)

func TestSummary(t *testing.T) {
	scanner := NewScanner(constants.Languages)
	tests := []struct {
		name       string
		scanResult []scanResult
		want       *Summary
	}{
		{
			name: "Should return the expected summary from scan result",
			scanResult: []scanResult{
				{
					Metadata: analyzer.FileMetadata{
						FilePath:  "/tmp/main.go",
						Extension: ".go",
						Language:  "Golang",
					},
					Lines:      167,
					CodeLines:  137,
					BlankLines: 14,
					Comments:   16,
				},
			},
			want: &Summary{
				Languages: map[string]*languageResult{
					"Golang": {
						Lines:      167,
						CodeLines:  137,
						BlankLines: 14,
						Comments:   16,
					},
				},
				Files: []fileResult{
					{
						Path:       "/tmp/main.go",
						Lines:      167,
						CodeLines:  137,
						BlankLines: 14,
						Comments:   16,
					},
				},
				FilesByLanguage: map[string]int{
					"Golang": 1,
				},
				TotalFiles:      1,
				TotalLines:      167,
				TotalCodeLines:  137,
				TotalBlankLines: 14,
				TotalComments:   16,
			},
		},
		{
			name: "Should return the expected summary from scan result with multiple files",
			scanResult: []scanResult{
				{
					Metadata: analyzer.FileMetadata{
						FilePath:  "/tmp/main.go",
						Extension: ".go",
						Language:  "Golang",
					},
					Lines:      167,
					CodeLines:  137,
					BlankLines: 14,
					Comments:   16,
				},
				{
					Metadata: analyzer.FileMetadata{
						FilePath:  "/tmp/test.go",
						Extension: ".go",
						Language:  "Golang",
					},
					Lines:      100,
					CodeLines:  80,
					BlankLines: 10,
					Comments:   2,
				},
				{
					Metadata: analyzer.FileMetadata{
						FilePath:  "/tmp/main.cpp",
						Extension: ".cpp",
						Language:  "CPP",
					},
					Lines:      100,
					CodeLines:  80,
					BlankLines: 10,
					Comments:   2,
				},
			},
			want: &Summary{
				Languages: map[string]*languageResult{
					"Golang": {
						Lines:      267,
						CodeLines:  217,
						BlankLines: 24,
						Comments:   18,
					},
					"CPP": {
						Lines:      100,
						CodeLines:  80,
						BlankLines: 10,
						Comments:   2,
					},
				},
				Files: []fileResult{
					{
						Path:       "/tmp/main.go",
						Lines:      167,
						CodeLines:  137,
						BlankLines: 14,
						Comments:   16,
					},
					{
						Path:       "/tmp/test.go",
						Lines:      100,
						CodeLines:  80,
						BlankLines: 10,
						Comments:   2,
					},
					{
						Path:       "/tmp/main.cpp",
						Lines:      100,
						CodeLines:  80,
						BlankLines: 10,
						Comments:   2,
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := scanner.Summary(tt.scanResult)
			require.Equal(t, tt.want, got)
		})
	}
}
