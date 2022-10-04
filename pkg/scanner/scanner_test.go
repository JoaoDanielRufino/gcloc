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
			Lines:      20,
			CodeLines:  15,
			BlankLines: 4,
			Comments:   1,
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
			Lines:      13,
			CodeLines:  8,
			BlankLines: 3,
			Comments:   2,
		},
		{
			Metadata:   files[5],
			Lines:      13,
			CodeLines:  8,
			BlankLines: 3,
			Comments:   2,
		},
	}

	result, err := scanner.Scan(files)
	require.NoError(t, err)

	require.Equal(t, areScanResultsEqual(expected, result), true)
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

func areScanResultsEqual(wanted []scanResult, files []scanResult) bool {
	if len(wanted) != len(files) {
		return false
	}

	globalPresent := true

	for _, want := range wanted {
		present := false
		for _, file := range files {
			if file == want {
				present = true
			}
		}
		globalPresent = (present == true)
	}

	return globalPresent
}
