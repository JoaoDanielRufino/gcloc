package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvertToMap(t *testing.T) {
	tests := []struct {
		name string
		arr  []string
		want map[string]bool
	}{
		{
			name: "Should convert array to map",
			arr:  []string{".go", ".c", ".cpp"},
			want: map[string]bool{
				".go":  true,
				".c":   true,
				".cpp": true,
			},
		},
		{
			name: "Should convert array(with repeated values) to map",
			arr:  []string{".go", ".c", ".c", ".cpp"},
			want: map[string]bool{
				".go":  true,
				".c":   true,
				".cpp": true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertToMap(tt.arr)
			require.Equal(t, tt.want, got)
		})
	}
}
