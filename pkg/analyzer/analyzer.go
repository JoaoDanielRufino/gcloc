package analyzer

import (
	"io/fs"
	"path/filepath"
	"strings"
)

type Analyzer struct {
	path                string
	excludePaths        []string
	supportedExtensions map[string]string
}

type FileMetadata struct {
	FilePath  string
	Extension string
	Language  string
}

func NewAnalyzer(path string, excludePaths []string, extensions map[string]string) *Analyzer {
	return &Analyzer{
		path:                path,
		excludePaths:        excludePaths,
		supportedExtensions: extensions,
	}
}

func (a *Analyzer) MatchingFiles() ([]FileMetadata, error) {
	var files []FileMetadata

	err := filepath.Walk(a.path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		fileExtension := a.getFileExtension(path)
		if a.canAdd(path, fileExtension) {
			fm := FileMetadata{
				FilePath:  path,
				Extension: fileExtension,
				Language:  a.supportedExtensions[fileExtension],
			}
			files = append(files, fm)
		}

		return nil
	})

	return files, err
}

func (a *Analyzer) getFileExtension(path string) string {
	extension := filepath.Ext(path)

	if extension == "" {
		extension = filepath.Base(path)
	}

	return extension
}

func (a *Analyzer) canAdd(path string, extension string) bool {
	for _, pathToExclude := range a.excludePaths {
		if strings.HasPrefix(path, pathToExclude) {
			return false
		}
	}

	_, ok := a.supportedExtensions[extension]
	return ok
}
