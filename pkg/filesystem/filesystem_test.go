package filesystem

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetExcludePaths(t *testing.T) {
	baseDir := filepath.Join("..", "..", "test", "fixtures", "filesystem_samples")

	tests := []struct {
		name              string
		path              string
		excludePathsParam []string
		want              []string
	}{
		{
			name:              "Should exclude all pkg dirs",
			path:              baseDir,
			excludePathsParam: []string{"pkg"},
			want: []string{
				filepath.Join(baseDir, "pkg"),
			},
		},
		{
			name:              "Should exclude only hello pkg",
			path:              baseDir,
			excludePathsParam: []string{"*/hello"},
			want: []string{
				filepath.Join(baseDir, "pkg", "hello"),
			},
		},
		{
			name:              "Should exclude all pkg .go files",
			path:              baseDir,
			excludePathsParam: []string{"pkg/*/*.go", "pkg/*.go"},
			want: []string{
				filepath.Join(baseDir, "pkg", "hello", "hello.go"),
				filepath.Join(baseDir, "pkg", "sum", "sum.go"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetExcludePaths(tt.path, tt.excludePathsParam)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
