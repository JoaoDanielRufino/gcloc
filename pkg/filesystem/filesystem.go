package filesystem

import (
	"path/filepath"
)

func GetExcludePaths(path string, excludePathsParam []string) ([]string, error) {
	var excludeDirs []string

	for _, dir := range excludePathsParam {
		matches, err := filepath.Glob(filepath.Join(path, dir))
		if err != nil {
			return []string{}, err
		}
		excludeDirs = append(excludeDirs, matches...)
	}

	return excludeDirs, nil
}
