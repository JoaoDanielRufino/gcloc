package sorter

import (
	"testing"

	"github.com/JoaoDanielRufino/gcloc/pkg/scanner"
	"github.com/stretchr/testify/require"
)

var summary = &scanner.Summary{
	Languages: map[string]*scanner.LanguageResult{
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
	Files: []scanner.FileResult{
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
}

func TestNewLanguageSorter(t *testing.T) {
	languageSorter := NewLanguageSorter("asc")
	expect := LanguageSorter{
		baseSorter{"ASC"},
	}
	require.Equal(t, expect, languageSorter)
}

func TestLangaugeSorterOrderByLanguage(t *testing.T) {
	tests := []struct {
		name           string
		languageSorter LanguageSorter
		summary        *scanner.Summary
		want           *SortedSummary
	}{
		{
			name:           "Should sort summary based on language on ascending order",
			languageSorter: NewLanguageSorter("asc"),
			summary:        summary,
			want: &SortedSummary{
				Results: []Result{
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
			name:           "Should sort summary based on language on descending order",
			languageSorter: NewLanguageSorter("desc"),
			summary:        summary,
			want: &SortedSummary{
				Results: []Result{
					{
						Name:       "Golang",
						Lines:      267,
						CodeLines:  217,
						BlankLines: 24,
						Comments:   18,
					},
					{
						Name:       "CPP",
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
			sortedSummary := tt.languageSorter.OrderByLanguage(tt.summary)
			require.Equal(t, tt.want, sortedSummary)
		})
	}
}

func TestLangaugeSorterOrderByCodeLines(t *testing.T) {
	tests := []struct {
		name           string
		languageSorter LanguageSorter
		summary        *scanner.Summary
		want           *SortedSummary
	}{
		{
			name:           "Should sort summary based on code lines on ascending order",
			languageSorter: NewLanguageSorter("asc"),
			summary:        summary,
			want: &SortedSummary{
				Results: []Result{
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
			name:           "Should sort summary based on code lines on descending order",
			languageSorter: NewLanguageSorter("desc"),
			summary:        summary,
			want: &SortedSummary{
				Results: []Result{
					{
						Name:       "Golang",
						Lines:      267,
						CodeLines:  217,
						BlankLines: 24,
						Comments:   18,
					},
					{
						Name:       "CPP",
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
			sortedSummary := tt.languageSorter.OrderByCodeLines(tt.summary)
			require.Equal(t, tt.want, sortedSummary)
		})
	}
}

func TestLangaugeSorterOrderByLines(t *testing.T) {
	tests := []struct {
		name           string
		languageSorter LanguageSorter
		summary        *scanner.Summary
		want           *SortedSummary
	}{
		{
			name:           "Should sort summary based on lines on ascending order",
			languageSorter: NewLanguageSorter("asc"),
			summary:        summary,
			want: &SortedSummary{
				Results: []Result{
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
			name:           "Should sort summary based on lines on descending order",
			languageSorter: NewLanguageSorter("desc"),
			summary:        summary,
			want: &SortedSummary{
				Results: []Result{
					{
						Name:       "Golang",
						Lines:      267,
						CodeLines:  217,
						BlankLines: 24,
						Comments:   18,
					},
					{
						Name:       "CPP",
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
			sortedSummary := tt.languageSorter.OrderByLines(tt.summary)
			require.Equal(t, tt.want, sortedSummary)
		})
	}
}

func TestLangaugeSorterOrderByComment(t *testing.T) {
	tests := []struct {
		name           string
		languageSorter LanguageSorter
		summary        *scanner.Summary
		want           *SortedSummary
	}{
		{
			name:           "Should sort summary based on comment on ascending order",
			languageSorter: NewLanguageSorter("asc"),
			summary:        summary,
			want: &SortedSummary{
				Results: []Result{
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
			name:           "Should sort summary based on comment on descending order",
			languageSorter: NewLanguageSorter("desc"),
			summary:        summary,
			want: &SortedSummary{
				Results: []Result{
					{
						Name:       "Golang",
						Lines:      267,
						CodeLines:  217,
						BlankLines: 24,
						Comments:   18,
					},
					{
						Name:       "CPP",
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
			sortedSummary := tt.languageSorter.OrderByComments(tt.summary)
			require.Equal(t, tt.want, sortedSummary)
		})
	}
}

func TestLangaugeSorterOrderByBlakLines(t *testing.T) {
	tests := []struct {
		name           string
		languageSorter LanguageSorter
		summary        *scanner.Summary
		want           *SortedSummary
	}{
		{
			name:           "Should sort summary based on blank lines on ascending order",
			languageSorter: NewLanguageSorter("asc"),
			summary:        summary,
			want: &SortedSummary{
				Results: []Result{
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
			name:           "Should sort summary based on blank lines on descending order",
			languageSorter: NewLanguageSorter("desc"),
			summary:        summary,
			want: &SortedSummary{
				Results: []Result{
					{
						Name:       "Golang",
						Lines:      267,
						CodeLines:  217,
						BlankLines: 24,
						Comments:   18,
					},
					{
						Name:       "CPP",
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
			sortedSummary := tt.languageSorter.OrderByBlankLines(tt.summary)
			require.Equal(t, tt.want, sortedSummary)
		})
	}
}

func TestLangaugeSorterOrderByFiles(t *testing.T) {
	tests := []struct {
		name           string
		languageSorter LanguageSorter
		summary        *scanner.Summary
		want           *SortedSummary
	}{
		{
			name:           "Should sort summary based on file count on ascending order",
			languageSorter: NewLanguageSorter("asc"),
			summary:        summary,
			want: &SortedSummary{
				Results: []Result{
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
			name:           "Should sort summary based on file count on descending order",
			languageSorter: NewLanguageSorter("desc"),
			summary:        summary,
			want: &SortedSummary{
				Results: []Result{
					{
						Name:       "Golang",
						Lines:      267,
						CodeLines:  217,
						BlankLines: 24,
						Comments:   18,
					},
					{
						Name:       "CPP",
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
			sortedSummary := tt.languageSorter.OrderByFiles(tt.summary)
			require.Equal(t, tt.want, sortedSummary)
		})
	}
}
