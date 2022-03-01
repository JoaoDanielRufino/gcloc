package analyzer

import (
	"io/fs"
	"path/filepath"
)

type Analyzer struct {
	Path                string
	ExcludeDirs         []string
	SupportedExtensions map[string]string
}

type fileMetadata struct {
	FilePath  string
	Extension string
	Language  string
}

func NewAnalyzer(path string, excludeDirs []string, extensions map[string]string) Analyzer {
	return Analyzer{
		Path:                path,
		ExcludeDirs:         excludeDirs,
		SupportedExtensions: extensions,
	}
}

func (a Analyzer) Analyze() ([]fileMetadata, error) {
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

func (a Analyzer) getFileExtension(path string) string {
	extension := filepath.Ext(path)

	if extension == "" {
		extension = filepath.Base(path)
	}

	return extension
}

func (a Analyzer) canAdd(path string, extension string) bool {
	if _, ok := a.SupportedExtensions[extension]; ok {
		return true
	}

	return false
}
