package analyzer

import (
	"io/fs"
	"path/filepath"
	"strings"
)

type Analyzer struct {
	Path                string
	ExcludePaths        []string
	SupportedExtensions map[string]string
}

type fileMetadata struct {
	FilePath  string
	Extension string
	Language  string
}

func NewAnalyzer(path string, excludePaths []string, extensions map[string]string) *Analyzer {
	return &Analyzer{
		Path:                path,
		ExcludePaths:        excludePaths,
		SupportedExtensions: extensions,
	}
}

func (a *Analyzer) Analyze() ([]fileMetadata, error) {
	var files []fileMetadata

	err := filepath.Walk(a.Path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		fileExtension := a.getFileExtension(path)
		if a.canAdd(path, fileExtension) {
			fm := fileMetadata{
				FilePath:  path,
				Extension: fileExtension,
				Language:  a.SupportedExtensions[fileExtension],
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
	for _, pathToExclude := range a.ExcludePaths {
		if strings.HasPrefix(path, pathToExclude) {
			return false
		}
	}

	if _, ok := a.SupportedExtensions[extension]; ok {
		return true
	}

	return false
}
