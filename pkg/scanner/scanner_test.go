package scanner

import (
	"path/filepath"
	"testing"

	"github.com/JoaoDanielRufino/gcloc/internal/constants"
	"github.com/JoaoDanielRufino/gcloc/pkg/analyzer"
	"github.com/stretchr/testify/require"
)

var languages = constants.Languages

func TestNewScanner(t *testing.T) {
	expected := &Scanner{SupportedLanguages: languages}
	scanner := NewScanner(languages)
	require.NotNil(t, scanner)
	require.Equal(t, expected, scanner)
}

func TestScan(t *testing.T) {
	codeSamplesDir := filepath.Join("..", "..", "test", "fixtures", "code_samples")
	scanner := NewScanner(languages)

	files := []analyzer.FileMetadata{
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
			FilePath:  filepath.Join(codeSamplesDir, "_main.java"),
			Extension: ".java",
			Language:  "Java",
		},
		{
			FilePath:  filepath.Join(codeSamplesDir, "index.html"),
			Extension: ".html",
			Language:  "HTML",
		},
		{
			FilePath:  filepath.Join(codeSamplesDir, "index.xml"),
			Extension: ".xml",
			Language:  "XML",
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
		{
			FilePath:  filepath.Join(codeSamplesDir, "Makefile"),
			Extension: "Makefile",
			Language:  "Makefile",
		},
		{
			FilePath:  filepath.Join(codeSamplesDir, "App.vue"),
			Extension: ".vue",
			Language:  "Vue",
		},
		{
			FilePath:  filepath.Join(codeSamplesDir, "readme.md"),
			Extension: ".md",
			Language:  "Markdown",
		},
		{
			FilePath:  filepath.Join(codeSamplesDir, "example.yml"),
			Extension: ".yml",
			Language:  "YAML",
		},
	}

	expected := []scanResult{
		{
			Metadata:   files[0],
			Lines:      13,
			CodeLines:  8,
			BlankLines: 3,
			Comments:   2,
		},
		{
			Metadata:   files[1],
			Lines:      17,
			CodeLines:  10,
			BlankLines: 5,
			Comments:   2,
		},
		{
			Metadata:   files[2],
			Lines:      24,
			CodeLines:  15,
			BlankLines: 5,
			Comments:   4,
		},
		{
			Metadata:   files[3],
			Lines:      167,
			CodeLines:  137,
			BlankLines: 14,
			Comments:   16,
		},
		{
			Metadata:   files[4],
			Lines:      40,
			CodeLines:  38,
			BlankLines: 0,
			Comments:   2,
		},
		{
			Metadata:   files[5],
			Lines:      20,
			CodeLines:  9,
			BlankLines: 5,
			Comments:   6,
		},
		{
			Metadata:   files[6],
			Lines:      13,
			CodeLines:  8,
			BlankLines: 3,
			Comments:   2,
		},
		{
			Metadata:   files[7],
			Lines:      13,
			CodeLines:  8,
			BlankLines: 3,
			Comments:   2,
		},
		{
			Metadata:   files[8],
			Lines:      23,
			CodeLines:  12,
			BlankLines: 4,
			Comments:   7,
		},
		{
			Metadata:   files[9],
			Lines:      30,
			CodeLines:  18,
			BlankLines: 4,
			Comments:   8,
		},
		{
			Metadata:   files[10],
			Lines:      15,
			CodeLines:  7,
			BlankLines: 3,
			Comments:   5,
		},
		{
			Metadata:   files[11],
			Lines:      14,
			CodeLines:  13,
			BlankLines: 1,
			Comments:   0,
		},
		{
			Metadata:   files[12],
			Lines:      6,
			CodeLines:  2,
			BlankLines: 1,
			Comments:   3,
		},
		{
			Metadata:   files[13],
			Lines:      24,
			CodeLines:  19,
			BlankLines: 4,
			Comments:   1,
		},
		{
			Metadata:   files[14],
			Lines:      51,
			CodeLines:  39,
			BlankLines: 7,
			Comments:   5,
		},
		{
			Metadata:   files[15],
			Lines:      24,
			CodeLines:  18,
			BlankLines: 6,
			Comments:   0,
		},
		{
			Metadata:   files[15],
			Lines:      28,
			CodeLines:  25,
			BlankLines: 1,
			Comments:   2,
		},
	}

	result, err := scanner.Scan(files)
	require.NoError(t, err)

	require.ElementsMatch(t, expected, result)
}
