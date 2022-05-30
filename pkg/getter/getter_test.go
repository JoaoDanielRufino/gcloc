package getter

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetter(t *testing.T) {
	testDir, _ := filepath.Abs(filepath.Join("..", "..", "test"))

	tests := []struct {
		name string
		src  string
		want string
	}{
		{
			name: "Should create symbolic link without errors",
			src:  filepath.Join("..", "..", "test", "fixtures", "code_samples"),
			want: filepath.Join(testDir, "fixtures", "code_samples"),
		},
		{
			name: "Should extract from github repo without errors",
			src:  "github.com/JoaoDanielRufino/gcloc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Getter(tt.src)
			require.NoError(t, err)

			if tt.want != "" {
				require.Equal(t, tt.want, got)
			}
		})
	}
}
