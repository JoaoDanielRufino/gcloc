package analyzer

import (
	"path/filepath"
	"testing"

	"github.com/JoaoDanielRufino/gcloc/internal/constants"
	"github.com/JoaoDanielRufino/gcloc/pkg/gcloc/language"
	"github.com/stretchr/testify/require"
)

var extensions = getExtensionsMap(constants.Languages)

func TestNewAnalyzer(t *testing.T) {
	expected := &Analyzer{
		SupportedExtensions: extensions,
		path:                "test/",
		excludePaths:        []string{"test"},
		excludeExtensions:   map[string]bool{".go": true},
		includeExtensions:   map[string]bool{},
	}
	analyser := NewAnalyzer(
		"test/",
		[]string{"test"},
		map[string]bool{".go": true},
		map[string]bool{},
		extensions,
	)
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
			name: "Should return matching files without errors",
			analyzer: NewAnalyzer(
				codeSamplesDir,
				[]string{},
				map[string]bool{},
				map[string]bool{},
				extensions,
			),
			want: []FileMetadata{
				{
					FilePath:  filepath.Join(codeSamplesDir, "_main.c"),
					Extension: ".c",
					Language:  "C",
				},
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
			name: "Should exclude files or dirs without errors",
			analyzer: NewAnalyzer(
				codeSamplesDir,
				[]string{filepath.Join(codeSamplesDir, "main.go")},
				map[string]bool{},
				map[string]bool{},
				extensions,
			),
			want: []FileMetadata{
				{
					FilePath:  filepath.Join(codeSamplesDir, "_main.c"),
					Extension: ".c",
					Language:  "C",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "index.html"),
					Extension: ".html",
					Language:  "HTML",
				},
			},
		},
		{
			name: "Should exclude extensions without errors",
			analyzer: NewAnalyzer(
				codeSamplesDir,
				[]string{},
				map[string]bool{".go": true},
				map[string]bool{},
				extensions,
			),
			want: []FileMetadata{
				{
					FilePath:  filepath.Join(codeSamplesDir, "_main.c"),
					Extension: ".c",
					Language:  "C",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "index.html"),
					Extension: ".html",
					Language:  "HTML",
				},
			},
		},
		{
			name: "Should include extensions without errors",
			analyzer: NewAnalyzer(
				codeSamplesDir,
				[]string{},
				map[string]bool{},
				map[string]bool{".go": true},
				extensions,
			),
			want: []FileMetadata{
				{
					FilePath:  filepath.Join(codeSamplesDir, "main.go"),
					Extension: ".go",
					Language:  "Golang",
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

func getExtensionsMap(languages language.Languages) map[string]string {
	extensions := map[string]string{}

	for language, languageInfo := range languages {
		for _, extension := range languageInfo.Extensions {
			extensions[extension] = language
		}
	}

	return extensions
}
