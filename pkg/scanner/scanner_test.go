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
	}

	expected := []scanResult{
		{
			Metadata:   files[0],
			Lines:      18,
			CodeLines:  8,
			BlankLines: 4,
			Comments:   6,
		},
		{
			Metadata:   files[1],
			Lines:      25,
			CodeLines:  10,
			BlankLines: 8,
			Comments:   7,
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
			Lines:      20,
			CodeLines:  9,
			BlankLines: 5,
			Comments:   6,
		},
		{
			Metadata:   files[5],
			Lines:      22,
			CodeLines:  8,
			BlankLines: 3,
			Comments:   2,
		},
		{
			Metadata:   files[6],
			Lines:      13,
			CodeLines:  8,
			BlankLines: 6,
			Comments:   8,
		},
		{
			Metadata:   files[7],
			Lines:      23,
			CodeLines:  12,
			BlankLines: 4,
			Comments:   7,
		},
		{
			Metadata:   files[8],
			Lines:      30,
			CodeLines:  18,
			BlankLines: 4,
			Comments:   8,
		},
		{
			Metadata:   files[9],
			Lines:      15,
			CodeLines:  7,
			BlankLines: 3,
			Comments:   5,
		},
	}

	result, err := scanner.Scan(files)
	require.NoError(t, err)

	require.ElementsMatch(t, expected, result)
}
