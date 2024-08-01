package junction

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestCreate(t *testing.T) {
	target, err := os.Getwd()
	if err != nil {
		t.Fatal(fmt.Errorf("os.Getwd: %w", err))
	}

	t.Logf("target is %q", target)
	mountpoint := filepath.Join(os.TempDir(), "junctionTest")
	t.Logf("mountpoint is %q", mountpoint)

	err = Create(target, mountpoint)
	if err != nil {
		t.Fatal(fmt.Errorf("Create: %w", err))
	}
	t.Log("mountpoint created")
	defer os.Remove(mountpoint)

	link, err := os.Readlink(mountpoint)
	if err != nil {
		t.Fatal(fmt.Errorf("os.Readlink: %w", err))
	}
	t.Log("hard link read")
	if target != link {
		t.Fatal(fmt.Errorf("os.Readlink: linked path differs : '%s' != '%s'", target, link))
	}

	file := filepath.Join(mountpoint, "junction.go")
	fd, err := os.Open(file)
	if err != nil {
		t.Fatal(fmt.Errorf("mount may be fake because '%s' can not be open", file))
	}
	t.Log("file checked")
	defer fd.Close()
}
