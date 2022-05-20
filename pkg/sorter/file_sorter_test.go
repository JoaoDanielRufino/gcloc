package sorter

import (
	"testing"

	"github.com/JoaoDanielRufino/gcloc/pkg/scanner"
	"github.com/stretchr/testify/require"
)

func TestNewFileSorter(t *testing.T) {
	fileSorter := NewFileSorter("asc")
	expect := FileSorter{
		baseSorter{"ASC"},
	}
	require.Equal(t, expect, fileSorter)
}

func TestFileSorterOrderByLanguage(t *testing.T) {
	tests := []struct {
		name       string
		fileSorter FileSorter
		summary    *scanner.Summary
		want       *SortedSummary
	}{
		{
			name:       "Should sort summary based on language on ascending order",
			fileSorter: NewFileSorter("asc"),
			summary:    summary,
			want: &SortedSummary{
				Results: []Result{
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
		{
			name:       "Should sort summary based on language on descending order",
			fileSorter: NewFileSorter("desc"),
			summary:    summary,
			want: &SortedSummary{
				Results: []Result{
					{
						Name:       "/tmp/test.go",
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
						Name:       "/tmp/main.cpp",
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
			sortedSummary := tt.fileSorter.OrderByLanguage(tt.summary)
			require.Equal(t, tt.want, sortedSummary)
		})
	}
}

func TestFileSorterOrderByCodeLines(t *testing.T) {
	tests := []struct {
		name       string
		fileSorter FileSorter
		summary    *scanner.Summary
		want       *SortedSummary
	}{
		{
			name:       "Should sort summary based on code lines on ascending order",
			fileSorter: NewFileSorter("asc"),
			summary:    summary,
			want: &SortedSummary{
				Results: []Result{
					{
						Name:       "/tmp/test.go",
						Lines:      100,
						CodeLines:  80,
						BlankLines: 10,
						Comments:   2,
					},
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
				},
				TotalLines:      367,
				TotalCodeLines:  297,
				TotalBlankLines: 34,
				TotalComments:   20,
			},
		},
		{
			name:       "Should sort summary based on code lines on descending order",
			fileSorter: NewFileSorter("desc"),
			summary:    summary,
			want: &SortedSummary{
				Results: []Result{
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
					{
						Name:       "/tmp/main.cpp",
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
			sortedSummary := tt.fileSorter.OrderByCodeLines(tt.summary)
			require.Equal(t, tt.want, sortedSummary)
		})
	}
}

func TestFileSorterOrderByLines(t *testing.T) {
	tests := []struct {
		name       string
		fileSorter FileSorter
		summary    *scanner.Summary
		want       *SortedSummary
	}{
		{
			name:       "Should sort summary based on lines on ascending order",
			fileSorter: NewFileSorter("asc"),
			summary:    summary,
			want: &SortedSummary{
				Results: []Result{
					{
						Name:       "/tmp/test.go",
						Lines:      100,
						CodeLines:  80,
						BlankLines: 10,
						Comments:   2,
					},
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
				},
				TotalLines:      367,
				TotalCodeLines:  297,
				TotalBlankLines: 34,
				TotalComments:   20,
			},
		},
		{
			name:       "Should sort summary based on lines on descending order",
			fileSorter: NewFileSorter("desc"),
			summary:    summary,
			want: &SortedSummary{
				Results: []Result{
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
					{
						Name:       "/tmp/main.cpp",
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
			sortedSummary := tt.fileSorter.OrderByLines(tt.summary)
			require.Equal(t, tt.want, sortedSummary)
		})
	}
}

func TestFileSorterOrderByComment(t *testing.T) {
	tests := []struct {
		name       string
		fileSorter FileSorter
		summary    *scanner.Summary
		want       *SortedSummary
	}{
		{
			name:       "Should sort summary based on comment on ascending order",
			fileSorter: NewFileSorter("asc"),
			summary:    summary,
			want: &SortedSummary{
				Results: []Result{
					{
						Name:       "/tmp/test.go",
						Lines:      100,
						CodeLines:  80,
						BlankLines: 10,
						Comments:   2,
					},
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
				},
				TotalLines:      367,
				TotalCodeLines:  297,
				TotalBlankLines: 34,
				TotalComments:   20,
			},
		},
		{
			name:       "Should sort summary based on comment on descending order",
			fileSorter: NewFileSorter("desc"),
			summary:    summary,
			want: &SortedSummary{
				Results: []Result{
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
					{
						Name:       "/tmp/main.cpp",
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
			sortedSummary := tt.fileSorter.OrderByComments(tt.summary)
			require.Equal(t, tt.want, sortedSummary)
		})
	}
}

func TestFileSorterOrderByBlakLines(t *testing.T) {
	tests := []struct {
		name       string
		fileSorter FileSorter
		summary    *scanner.Summary
		want       *SortedSummary
	}{
		{
			name:       "Should sort summary based on blank lines on ascending order",
			fileSorter: NewFileSorter("asc"),
			summary:    summary,
			want: &SortedSummary{
				Results: []Result{
					{
						Name:       "/tmp/test.go",
						Lines:      100,
						CodeLines:  80,
						BlankLines: 10,
						Comments:   2,
					},
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
				},
				TotalLines:      367,
				TotalCodeLines:  297,
				TotalBlankLines: 34,
				TotalComments:   20,
			},
		},
		{
			name:       "Should sort summary based on blank lines on descending order",
			fileSorter: NewFileSorter("desc"),
			summary:    summary,
			want: &SortedSummary{
				Results: []Result{
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
					{
						Name:       "/tmp/main.cpp",
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
			sortedSummary := tt.fileSorter.OrderByBlankLines(tt.summary)
			require.Equal(t, tt.want, sortedSummary)
		})
	}
}
