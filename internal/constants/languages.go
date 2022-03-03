package constants

import (
	"github.com/JoaoDanielRufino/gcloc/pkg/gcloc/language"
)

var Languages = map[string]language.LanguageInfo{
	"C++": {
		LineComments:     []string{"//"},
		MultiLineComment: [][]string{{"/*", "*/"}},
	},
	"Golang": {
		LineComments:     []string{"//"},
		MultiLineComment: [][]string{{"/*", "*/"}},
	},
	"JavaScript": {
		LineComments:     []string{"//"},
		MultiLineComment: [][]string{{"/*", "*/"}},
	},
	"TypeScript": {
		LineComments:     []string{"//"},
		MultiLineComment: [][]string{{"/*", "*/"}},
	},
}
