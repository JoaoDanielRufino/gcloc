package filesystem

import (
	"path/filepath"
)

func GetExcludeFiles(path string, excludeDirParams []string) ([]string, error) {
	var excludeDirs []string

	for _, dir := range excludeDirParams {
		matches, err := filepath.Glob(filepath.Join(path, dir))
		if err != nil {
			return []string{}, err
		}
		excludeDirs = append(excludeDirs, matches...)
	}

	return excludeDirs, nil
}
