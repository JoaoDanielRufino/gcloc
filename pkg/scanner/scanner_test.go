package scanner

import (
	"path/filepath"
	"testing"

	"github.com/JoaoDanielRufino/gcloc/internal/constants"
	"github.com/JoaoDanielRufino/gcloc/pkg/analyzer"
	"github.com/JoaoDanielRufino/gcloc/pkg/gcloc/language"
	"github.com/stretchr/testify/require"
)

var languages = constants.Languages
var extensions = getExtensionsMap(languages)

func TestNewScanner(t *testing.T) {
	expected := &Scanner{SupportedLanguages: languages}
	scanner := NewScanner(languages)
	require.NotNil(t, scanner)
	require.Equal(t, expected, scanner)
}

func TestScan(t *testing.T) {
	codeSamplesDir := filepath.Join("..", "..", "test", "fixtures", "code_samples")
	fileAnalyzer := analyzer.NewAnalyzer(codeSamplesDir, []string{}, map[string]bool{}, map[string]bool{}, extensions)
	scanner := NewScanner(languages)

	files, _ := fileAnalyzer.MatchingFiles()

	expected := []scanResult{
		{
			Metadata:   files[0],
			Lines:      167,
			CodeLines:  137,
			BlankLines: 14,
			Comments:   16,
		},
		{
			Metadata:   files[1],
			Lines:      20,
			CodeLines:  9,
			BlankLines: 5,
			Comments:   6,
		},
	}

	result, err := scanner.Scan(files)
	require.NoError(t, err)
	require.Equal(t, expected, result)
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
