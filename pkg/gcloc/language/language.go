package language

type LanguageInfo struct {
	LineComments      []string
	MultiLineComments [][]string
}

type Languages map[string]LanguageInfo
