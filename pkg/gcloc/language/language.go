package language

type LanguageInfo struct {
	LineComments      []string
	MultiLineComments [][]string
	Extensions        []string
}

type Languages map[string]LanguageInfo
