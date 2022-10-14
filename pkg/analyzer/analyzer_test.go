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
		excludeExtensions:   map[string]struct{}{".go": {}},
		includeExtensions:   map[string]struct{}{},
	}
	analyser := NewAnalyzer(
		"test/",
		[]string{"test"},
		map[string]struct{}{".go": {}},
		map[string]struct{}{},
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
				map[string]struct{}{},
				map[string]struct{}{},
				extensions,
			),
			want: []FileMetadata{
				{
					FilePath:  filepath.Join(codeSamplesDir, "_main.c"),
					Extension: ".c",
					Language:  "C",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "_main.cc"),
					Extension: ".cc",
					Language:  "C++",
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
				{
					FilePath:  filepath.Join(codeSamplesDir, "main.js"),
					Extension: ".js",
					Language:  "JavaScript",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "main.ts"),
					Extension: ".ts",
					Language:  "TypeScript",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "_main.java"),
					Extension: ".java",
					Language:  "Java",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "main.py"),
					Extension: ".py",
					Language:  "Python",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "main.rb"),
					Extension: ".rb",
					Language:  "Ruby",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "Main.kt"),
					Extension: ".kt",
					Language:  "Kotlin",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "hello.json"),
					Extension: ".json",
					Language:  "Json",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "main.php"),
					Extension: ".php",
					Language:  "PHP",
				},
			},
		},
		{
			name: "Should exclude files or dirs without errors",
			analyzer: NewAnalyzer(
				codeSamplesDir,
				[]string{filepath.Join(codeSamplesDir, "main.go")},
				map[string]struct{}{},
				map[string]struct{}{},
				extensions,
			),
			want: []FileMetadata{
				{
					FilePath:  filepath.Join(codeSamplesDir, "_main.c"),
					Extension: ".c",
					Language:  "C",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "_main.cc"),
					Extension: ".cc",
					Language:  "C++",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "index.html"),
					Extension: ".html",
					Language:  "HTML",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "main.js"),
					Extension: ".js",
					Language:  "JavaScript",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "main.ts"),
					Extension: ".ts",
					Language:  "TypeScript",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "_main.java"),
					Extension: ".java",
					Language:  "Java",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "main.py"),
					Extension: ".py",
					Language:  "Python",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "main.rb"),
					Extension: ".rb",
					Language:  "Ruby",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "Main.kt"),
					Extension: ".kt",
					Language:  "Kotlin",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "hello.json"),
					Extension: ".json",
					Language:  "Json",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "main.php"),
					Extension: ".php",
					Language:  "PHP",
				},
			},
		},
		{
			name: "Should exclude extensions without errors",
			analyzer: NewAnalyzer(
				codeSamplesDir,
				[]string{},
				map[string]struct{}{".go": {}},
				map[string]struct{}{},
				extensions,
			),
			want: []FileMetadata{
				{
					FilePath:  filepath.Join(codeSamplesDir, "_main.c"),
					Extension: ".c",
					Language:  "C",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "_main.cc"),
					Extension: ".cc",
					Language:  "C++",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "index.html"),
					Extension: ".html",
					Language:  "HTML",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "main.js"),
					Extension: ".js",
					Language:  "JavaScript",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "main.ts"),
					Extension: ".ts",
					Language:  "TypeScript",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "_main.java"),
					Extension: ".java",
					Language:  "Java",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "main.py"),
					Extension: ".py",
					Language:  "Python",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "main.rb"),
					Extension: ".rb",
					Language:  "Ruby",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "Main.kt"),
					Extension: ".kt",
					Language:  "Kotlin",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "hello.json"),
					Extension: ".json",
					Language:  "Json",
				},
				{
					FilePath:  filepath.Join(codeSamplesDir, "main.php"),
					Extension: ".php",
					Language:  "PHP",
				},
			},
		},
		{
			name: "Should include extensions without errors",
			analyzer: NewAnalyzer(
				codeSamplesDir,
				[]string{},
				map[string]struct{}{},
				map[string]struct{}{".go": {}},
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

			require.ElementsMatch(t, tt.want, files)
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
