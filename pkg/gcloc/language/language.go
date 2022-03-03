package language

type LanguageInfo struct {
	LineComments     []string
	MultiLineComment [][]string
}

type Languages map[string]LanguageInfo
