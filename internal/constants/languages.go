package constants

import "github.com/JoaoDanielRufino/gcloc/pkg/gcloc/language"

var Languages = language.Languages{
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
	"C Header": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".h"},
	},
	"C++": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".cpp", ".cc"},
	},
	"C++ Header": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".hh", ".hpp"},
	},
	"Golang": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".go"},
	},
	"HTML": {
		LineComments:      []string{},
		MultiLineComments: [][]string{{"<!--", "-->"}},
		Extensions:        []string{".html"},
	},
	"Haskell": {
		LineComments:      []string{"--"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".hs"},
	},
	"Java": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".java"},
	},
	"JavaScript": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".js", ".jsx"},
	},
	"PHP": {
		LineComments:      []string{"//", "#"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".php"},
	},
	"Perl": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{{"=", "=cut"}},
		Extensions:        []string{".pl"},
	},
	"Python": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{{"\"\"\"", "\"\"\""}},
		Extensions:        []string{".py"},
	},
	"TypeScript": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".ts", ".tsx"},
	},
	"Makefile": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{},
		Extensions:        []string{"Makefile"},
	},
	"YAML": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".yaml", ".yml"},
	},
}
