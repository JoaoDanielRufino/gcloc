package scanner

import (
	"bufio"
	"os"
	"strings"

	"github.com/JoaoDanielRufino/gcloc/pkg/analyzer"
	"github.com/JoaoDanielRufino/gcloc/pkg/gcloc/language"
)

type Scanner struct {
	supportedLanguages language.Languages
}

type scanResult struct {
	Metadata   analyzer.FileMetadata
	CodeLines  int
	BlankLines int
	Comments   int
}

func NewScanner(languages language.Languages) *Scanner {
	return &Scanner{
		supportedLanguages: languages,
	}
}

func (sc *Scanner) Scan(files []analyzer.FileMetadata) ([]scanResult, error) {
	var results []scanResult

	for _, file := range files {
		codeLines, blankLines, comments, err := sc.scanFile(file)
		if err != nil {
			return results, err
		}

		result := scanResult{
			Metadata:   file,
			CodeLines:  codeLines,
			BlankLines: blankLines,
			Comments:   comments,
		}
		results = append(results, result)
	}

	return results, nil
}

func (sc *Scanner) ChangeLanguages(languages language.Languages) {
	sc.supportedLanguages = languages
}

func (sc *Scanner) scanFile(file analyzer.FileMetadata) (int, int, int, error) {
	lines := 0
	blankLines := 0
	comments := 0

	f, err := os.Open(file.FilePath)
	if err != nil {
		return lines, blankLines, comments, err
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())

		lines++

		if len(line) == 0 {
			blankLines++
		}

		if sc.hasSingleLineComment(file, line) {
			comments++
		}
	}

	return lines, blankLines, comments, fileScanner.Err()
}

func (sc *Scanner) hasSingleLineComment(file analyzer.FileMetadata, line string) bool {
	lineComments := sc.supportedLanguages[file.Language].LineComments

	for _, lineComment := range lineComments {
		if strings.HasPrefix(line, lineComment) {
			return true
		}
	}

	return false
}
