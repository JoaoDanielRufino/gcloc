package scanner

import (
	"bufio"
	"os"
	"strings"

	"github.com/JoaoDanielRufino/gcloc/pkg/analyzer"
	"github.com/JoaoDanielRufino/gcloc/pkg/gcloc/language"
	"github.com/schollz/progressbar/v3"
)

type Scanner struct {
	SupportedLanguages language.Languages
}

type scanResult struct {
	Metadata   analyzer.FileMetadata
	Lines      int
	CodeLines  int
	BlankLines int
	Comments   int
}

func NewScanner(languages language.Languages) *Scanner {
	return &Scanner{
		SupportedLanguages: languages,
	}
}

func (sc *Scanner) Scan(files []analyzer.FileMetadata) ([]scanResult, error) {
	var results []scanResult
	progress := progressbar.NewOptions(
		len(files),
		progressbar.OptionSetDescription("Scanning files..."),
		progressbar.OptionShowBytes(true),
		progressbar.OptionShowCount(),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "=",
			SaucerHead:    ">",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)

	for _, file := range files {
		result, err := sc.scanFile(file)
		if err != nil {
			return results, err
		}
		progress.Add(1)
		results = append(results, result)
	}

	return results, progress.Clear()
}

func (sc *Scanner) scanFile(file analyzer.FileMetadata) (scanResult, error) {
	result := scanResult{Metadata: file}
	isInBlockComment := false
	var closeBlockCommentToken string

	f, err := os.Open(file.FilePath)
	if err != nil {
		return result, err
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	buffer := make([]byte, 128*1024)
	fileScanner.Buffer(buffer, 4096*1024)
	for fileScanner.Scan() {
		line := strings.TrimSpace(fileScanner.Text())

		if isInBlockComment {
			result.Comments++
			if sc.hasSecondMultiLineComment(line, closeBlockCommentToken) {
				isInBlockComment = false
			}
			continue
		}

		if sc.isBlankLine(line) {
			result.BlankLines++
			continue
		}

		if ok, secondCommentToken := sc.hasFirstMultiLineComment(file, line); ok {
			isInBlockComment = true
			closeBlockCommentToken = secondCommentToken
			result.Comments++
			if sc.hasSecondMultiLineComment(line, closeBlockCommentToken) {
				isInBlockComment = false
			}
			continue
		}

		if sc.hasSingleLineComment(file, line) {
			result.Comments++
			continue
		}

		result.CodeLines++
	}

	result.Lines = result.CodeLines + result.BlankLines + result.Comments

	return result, fileScanner.Err()
}

func (sc *Scanner) hasFirstMultiLineComment(file analyzer.FileMetadata, line string) (bool, string) {
	multiLineComments := sc.SupportedLanguages[file.Language].MultiLineComments

	for _, multiLineComment := range multiLineComments {
		firstCommentToken := multiLineComment[0]
		if strings.HasPrefix(line, firstCommentToken) {
			return true, multiLineComment[1]
		}
	}

	return false, ""
}

func (sc *Scanner) hasSecondMultiLineComment(line, commentToken string) bool {
	return strings.Contains(line, commentToken)
}

func (sc *Scanner) hasSingleLineComment(file analyzer.FileMetadata, line string) bool {
	lineComments := sc.SupportedLanguages[file.Language].LineComments

	for _, lineComment := range lineComments {
		if strings.HasPrefix(line, lineComment) {
			return true
		}
	}

	return false
}

func (sc *Scanner) isBlankLine(line string) bool {
	return len(line) == 0
}
