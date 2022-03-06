package analyzer

import (
	"path/filepath"
	"testing"

	"github.com/JoaoDanielRufino/gcloc/internal/constants"
	"github.com/stretchr/testify/require"
)

var extensions = constants.Extensions

func TestNewAnalyzer(t *testing.T) {
	expected := &Analyzer{
		SupportedExtensions: extensions,
		path:                "test/",
		excludePaths:        []string{"test"},
	}
	analyser := NewAnalyzer("test/", []string{"test"}, extensions)
	require.NotNil(t, analyser)
	require.Equal(t, expected, analyser)
}

func TestMatchingFiles(t *testing.T) {
	codeSamplesDir := filepath.Join("..", "..", "test", "fixtures", "code_samples")
	tests := []struct {
		name     string
		analyzer *Analyzer
		want     []FileMetadata
	}{
		{
			name:     "Should return matching files without errors",
			analyzer: NewAnalyzer(codeSamplesDir, []string{}, extensions),
			want: []FileMetadata{
				{
					FilePath:  filepath.Join(codeSamplesDir, "index.html"),
					Extension: ".html",
					Language:  "HTML",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "main.go"),
					Extension: ".go",
					Language:  "Golang",
				},
			},
		},
		{
			name:     "Should exclude files or dirs without errors",
			analyzer: NewAnalyzer(codeSamplesDir, []string{filepath.Join(codeSamplesDir, "main.go")}, extensions),
			want: []FileMetadata{
				{
					FilePath:  filepath.Join(codeSamplesDir, "index.html"),
					Extension: ".html",
					Language:  "HTML",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			files, err := tt.analyzer.MatchingFiles()
			require.NoError(t, err)
			require.Equal(t, tt.want, files)
		})
	}
}
