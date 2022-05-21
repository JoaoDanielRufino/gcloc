package gcloc

import (
	"path/filepath"
	"testing"

	"github.com/JoaoDanielRufino/gcloc/internal/constants"
	"github.com/JoaoDanielRufino/gcloc/pkg/analyzer"
	"github.com/JoaoDanielRufino/gcloc/pkg/gcloc/language"
	"github.com/JoaoDanielRufino/gcloc/pkg/reporter/prompt"
	"github.com/JoaoDanielRufino/gcloc/pkg/scanner"
	"github.com/JoaoDanielRufino/gcloc/pkg/sorter"
	"github.com/stretchr/testify/require"
)

func TestNewGCloc(t *testing.T) {
	tests := []struct {
		name      string
		params    Params
		languages language.Languages
		want      *GCloc
	}{
		{
			name: "Should create gcloc without errors",
			params: Params{
				Path:  filepath.Join("..", "..", "test", "fixtures", "code_samples"),
				Order: "ASC",
			},
			languages: constants.Languages,
			want: &GCloc{
				params: Params{
					Path:  filepath.Join("..", "..", "test", "fixtures", "code_samples"),
					Order: "ASC",
				},
				analyzer: analyzer.NewAnalyzer(
					filepath.Join("..", "..", "test", "fixtures", "code_samples"),
					nil,
					map[string]bool{},
					map[string]bool{},
					getExtensionsMap(constants.Languages),
				),
				scanner:  scanner.NewScanner(constants.Languages),
				sorter:   sorter.NewLanguageSorter("ASC"),
				reporter: prompt.PromptReporter{},
			},
		},
		{
			name: "Should create gcloc without errors with exclude extensions",
			params: Params{
				Path:              filepath.Join("..", "..", "test", "fixtures", "code_samples"),
				Order:             "ASC",
				ByFile:            true,
				ExcludeExtensions: []string{".go"},
			},
			languages: constants.Languages,
			want: &GCloc{
				params: Params{
					Path:              filepath.Join("..", "..", "test", "fixtures", "code_samples"),
					Order:             "ASC",
					ByFile:            true,
					ExcludeExtensions: []string{".go"},
				},
				analyzer: analyzer.NewAnalyzer(
					filepath.Join("..", "..", "test", "fixtures", "code_samples"),
					nil,
					map[string]bool{".go": true},
					map[string]bool{},
					getExtensionsMap(constants.Languages),
				),
				scanner:  scanner.NewScanner(constants.Languages),
				sorter:   sorter.NewFileSorter("ASC"),
				reporter: prompt.PromptReporter{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGCloc(tt.params, tt.languages)
			require.NoError(t, err)
			require.NotNil(t, got)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestRun(t *testing.T) {
	tests := []struct {
		name      string
		params    Params
		languages language.Languages
	}{
		{
			name: "Should run gcloc without errors",
			params: Params{
				Path:  filepath.Join("..", "..", "test", "fixtures", "code_samples"),
				Order: "ASC",
			},
			languages: constants.Languages,
		},
		{
			name: "Should run gcloc by file without errors",
			params: Params{
				Path:   filepath.Join("..", "..", "test", "fixtures", "code_samples"),
				Order:  "ASC",
				ByFile: true,
			},
			languages: constants.Languages,
		},
		{
			name: "Should run gcloc ordered by code without errors",
			params: Params{
				Path:        filepath.Join("..", "..", "test", "fixtures", "code_samples"),
				Order:       "ASC",
				OrderByCode: true,
			},
			languages: constants.Languages,
		},
		{
			name: "Should run gcloc ordered by lang without errors",
			params: Params{
				Path:        filepath.Join("..", "..", "test", "fixtures", "code_samples"),
				Order:       "ASC",
				OrderByLang: true,
			},
			languages: constants.Languages,
		},
		{
			name: "Should run gcloc ordered by line without errors",
			params: Params{
				Path:        filepath.Join("..", "..", "test", "fixtures", "code_samples"),
				Order:       "ASC",
				OrderByLine: true,
			},
			languages: constants.Languages,
		},
		{
			name: "Should run gcloc ordered by comment without errors",
			params: Params{
				Path:           filepath.Join("..", "..", "test", "fixtures", "code_samples"),
				Order:          "ASC",
				OrderByComment: true,
			},
			languages: constants.Languages,
		},
		{
			name: "Should run gcloc ordered by blank lines without errors",
			params: Params{
				Path:         filepath.Join("..", "..", "test", "fixtures", "code_samples"),
				Order:        "ASC",
				OrderByBlank: true,
			},
			languages: constants.Languages,
		},
		{
			name: "Should run gcloc ordered by file without errors",
			params: Params{
				Path:        filepath.Join("..", "..", "test", "fixtures", "code_samples"),
				Order:       "ASC",
				OrderByFile: true,
			},
			languages: constants.Languages,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gcloc, _ := NewGCloc(tt.params, tt.languages)
			err := gcloc.Run()
			require.NoError(t, err)
		})
	}
}

func TestChangeLanguages(t *testing.T) {
	gcloc, _ := NewGCloc(Params{Path: "/tmp"}, constants.Languages)

	newLanguages := language.Languages{
		"Assembly": {
			LineComments:      []string{"//", ";", "#", "@", "|", "!"},
			MultiLineComments: [][]string{{"/*", "*/"}},
			Extensions:        []string{".asm"},
		},
		"C": {
			LineComments:      []string{"//"},
			MultiLineComments: [][]string{{"/*", "*/"}},
			Extensions:        []string{".c"},
		},
	}

	extensions := map[string]string{
		".asm": "Assembly",
		".c":   "C",
	}

	gcloc.ChangeLanguages(newLanguages)
	require.Equal(t, newLanguages, gcloc.scanner.SupportedLanguages)
	require.Equal(t, extensions, gcloc.analyzer.SupportedExtensions)
}
