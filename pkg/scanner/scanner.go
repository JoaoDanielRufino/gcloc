package scanner

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/JoaoDanielRufino/gcloc/pkg/analyzer"
)

type Scanner struct {
	BufferSize int
}

type scanResult struct {
	Metadata  analyzer.FileMetadata
	CodeLines int
	Comments  int
}

func NewScanner() *Scanner {
	return &Scanner{
		BufferSize: 32,
	}
}

func (sc *Scanner) Scan(files []analyzer.FileMetadata) ([]scanResult, error) {
	var results []scanResult

	for _, file := range files {
		f, err := os.Open(file.FilePath)
		if err != nil {
			return results, err
		}

		codeLines, comments, err := sc.countLines(f)
		if err != nil {
			return results, err
		}

		result := scanResult{
			Metadata:  file,
			CodeLines: codeLines,
			Comments:  comments,
		}
		results = append(results, result)
		f.Close()
	}

	return results, nil
}

func (sc *Scanner) countLines(r io.Reader) (int, int, error) {
	fileScanner := bufio.NewScanner(r)
	lines := 0
	comments := 0

	for fileScanner.Scan() {
		line := strings.TrimLeft(fileScanner.Text(), " ")

		lines++

		if strings.HasPrefix(line, "//") {
			comments++
		}
	}

	err := fileScanner.Err()

	return lines, comments, err
}
