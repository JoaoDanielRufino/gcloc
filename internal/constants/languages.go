package constants

import "github.com/JoaoDanielRufino/gcloc/pkg/gcloc/language"

var Languages = language.Languages{
	"ActionScript": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".as"},
	},
	"Assembly": {
		LineComments:      []string{"//", ";", "#", "@", "|", "!"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".asm"},
	},
	"Batch": {
		LineComments:      []string{"REM", "rem"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".bat"},
	},
	"Bash": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".bash", ".sh"},
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
	"COBOL": {
		LineComments:      []string{"*", "/"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".cbl"},
	},
	"CoffeeScript": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{{"###", "###"}},
		Extensions:        []string{".coffee"},
	},
	"Clojure": {
		LineComments:      []string{";", ";;", "#_"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".clj"},
	},
	"C#": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".cs"},
	},
	"CSS": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".css"},
	},
	"Dart": {
		LineComments:      []string{"//", "///"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".dart"},
	},
	"Elixir": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".ex", ".exs"},
	},
	"Erlang": {
		LineComments:      []string{"%", "%%"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".erl"},
	},
	"Lua": {
		LineComments:      []string{"--"},
		MultiLineComments: [][]string{{"--[[", "]]"}},
		Extensions:        []string{".lua"},
	},
	"Lisp": {
		LineComments:      []string{";"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".lsp", ".lisp"},
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
	"Jupyter Notebook": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".ipynb"},
	},
	"Json": {
		LineComments:      []string{},
		MultiLineComments: [][]string{},
		Extensions:        []string{".json"},
	},
	"Kotlin": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".kt", ".kts"},
	},
	"Markdown": {
		LineComments:      []string{},
		MultiLineComments: [][]string{},
		Extensions:        []string{".md", ".markdown"},
	},
	"Maven": {
		LineComments:      []string{"<!--"},
		MultiLineComments: [][]string{{"<!--", "-->"}},
		Extensions:        []string{".maven"},
	},
	"Makefile": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{},
		Extensions:        []string{"Makefile"},
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
	"Plain Text": {
		LineComments:      []string{},
		MultiLineComments: [][]string{},
		Extensions:        []string{".txt", ".text"},
	},
	"PowerShell": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".ps1"},
	},
	"Protocol Buffers": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".proto"},
	},
	"Processing": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".pde"},
	},
	"Python": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{{"\"\"\"", "\"\"\""}},
		Extensions:        []string{".py"},
	},
	"R": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".r", ".R"},
	},
	"Rego": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".rego"},
	},
	"Ruby": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{{"=begin", "=end"}},
		Extensions:        []string{".rb"},
	},
	"Rust": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".rs"},
	},
	"Scala": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".scala"},
	},
	"Scss": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".scss"},
	},
	"SQL": {
		LineComments:      []string{"--"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".sql"},
	},
	"Swift": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".swift"},
	},
	"TypeScript": {
		LineComments:      []string{"//"},
		MultiLineComments: [][]string{{"/*", "*/"}},
		Extensions:        []string{".ts", ".tsx"},
	},
	"Vim": {
		LineComments:      []string{`"`},
		MultiLineComments: [][]string{},
		Extensions:        []string{".vim"},
	},
	"Vue": {
		LineComments:      []string{"<!--"},
		MultiLineComments: [][]string{{"<!--", "-->"}},
		Extensions:        []string{".vue"},
	},
	"XML": {
		LineComments:      []string{"<!--"},
		MultiLineComments: [][]string{{"<!--", "-->"}},
		Extensions:        []string{".xml", ".XML"},
	},
	"YAML": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".yaml", ".yml"},
	},
	"Zsh": {
		LineComments:      []string{"#"},
		MultiLineComments: [][]string{},
		Extensions:        []string{".zsh"},
	},
}
