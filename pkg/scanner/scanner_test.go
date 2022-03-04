package scanner

import (
	"path/filepath"
	"testing"

	"github.com/JoaoDanielRufino/gcloc/pkg/analyzer"
	"github.com/JoaoDanielRufino/gcloc/pkg/gcloc/language"
	"github.com/stretchr/testify/require"
)

var extensions = map[string]string{
	".cpp": "C++",
	".go":  "Golang",
	".js":  "JavaScript",
	".ts":  "TypeScript",
}

var languages = language.Languages{
	"C++": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
	},
	"Golang": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
	},
	"JavaScript": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
	},
	"TypeScript": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
	},
}

func TestNewScanner(t *testing.T) {
	scanner := NewScanner(languages)
	require.NotNil(t, scanner)
	require.Equal(t, languages, scanner.SupportedLanguages)
}

func TestScan(t *testing.T) {
	fixtureScannerDir := filepath.Join("..", "..", "test", "fixtures", "scanner")
	fileAnalyzer := analyzer.NewAnalyzer(fixtureScannerDir, []string{}, extensions)
	scanner := NewScanner(languages)

	files, _ := fileAnalyzer.MatchingFiles()

	expected := []scanResult{
		{
			Metadata:   files[0],
			CodeLines:  20,
			Comments:   6,
			BlankLines: 5,
		},
	}

	result, err := scanner.Scan(files)
	require.NoError(t, err)
	require.Equal(t, expected, result)
}
