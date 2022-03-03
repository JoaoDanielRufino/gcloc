package scanner

import (
	"bytes"
	"io"
	"os"

	"github.com/JoaoDanielRufino/gcloc/pkg/analyzer"
)

type Scanner struct {
	BufferSize int
}

type scanResult struct {
	Metadata  analyzer.FileMetadata
	CodeLines int
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

		codeLines, err := sc.countLines(f)
		if err != nil {
			return results, err
		}

		result := scanResult{
			Metadata:  file,
			CodeLines: codeLines,
		}
		results = append(results, result)
		f.Close()
	}

	return results, nil
}

func (sc *Scanner) countLines(r io.Reader) (int, error) {
	buffer := make([]byte, sc.BufferSize*1024)
	lineStep := []byte{'\n'}
	count := 0

	for {
		n, err := r.Read(buffer)
		count += bytes.Count(buffer[:n], lineStep)

		switch err {
		case io.EOF:
			return count, nil
		case nil:
			continue
		default:
			return count, err
		}
	}
}
