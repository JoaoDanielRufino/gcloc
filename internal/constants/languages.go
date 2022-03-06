package constants

import "github.com/JoaoDanielRufino/gcloc/pkg/gcloc/language"

var Languages = language.Languages{
	"Assembly": {
		LineComments:      []string{"//", ";", "#", "@", "|", "!"},
		MultiLineComments: [][]string{{"/*", "*/"}},
	},
	"C": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
	},
	"C Header": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
	},
	"C++": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
	},
	"C++ Header": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
	},
	"Golang": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
	},
	"HTML": {
		LineComments:      []string{},
		MultiLineComments: [][]string{{"<!--", "-->"}},
	},
	"Haskell": {
		LineComments:      []string{"--"},
		MultiLineComments: [][]string{},
	},
	"Java": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
	},
	"JavaScript": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
	},
	"PHP": {
		LineComments:      []string{"//", "#"},
		MultiLineComments: [][]string{{"/*", "*/"}},
	},
	"Prolog": {
		LineComments:      []string{"%"},
		MultiLineComments: [][]string{},
	},
	"Python": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{{"\"\"\"", "\"\"\""}},
	},
	"TypeScript": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
	},
	"Makefile": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{},
	},
	"YAML": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{},
	},
}
