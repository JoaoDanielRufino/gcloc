package getter

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	getter "github.com/hashicorp/go-getter"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func Getter(src string) (string, error) {
	dst := filepath.Join(os.TempDir(), fmt.Sprintf("gcloc-extract-%d", seededRand.Int()))

	client := &getter.Client{
		Src:  src,
		Dst:  dst,
		Mode: getter.ClientModeAny,
	}

	if err := client.Get(); err != nil {
		return "", err
	}

	info, err := os.Lstat(dst)
	if err != nil {
		return "", err
	}

	if info.Mode()&os.ModeSymlink != 0 {
		origin, err := os.Readlink(dst)
		if err != nil {
			return "", err
		}

		return origin, nil
	}

	return dst, nil
}
