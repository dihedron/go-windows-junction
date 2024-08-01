//go:build !windows
// +build !windows

package junction

import (
	"os"
)

func create(target, link string) error {
	return os.Symlink(target, link)
}
