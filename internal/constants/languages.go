package constants

import "github.com/JoaoDanielRufino/gcloc/pkg/gcloc"

var Languages = map[string]gcloc.LanguageInfo{
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
