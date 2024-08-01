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
	name := filepath.Join(os.TempDir(), "junctionTest")
	t.Logf("link is %q", name)

	err = Create(target, name)
	if err != nil {
		t.Fatal(fmt.Errorf("Create: %w", err))
	}
	t.Log("link created")
	defer os.Remove(name)

	link, err := os.Readlink(name)
	if err != nil {
		t.Fatal(fmt.Errorf("os.Readlink: %w", err))
	}
	t.Log("hard link read")
	if target != link {
		t.Fatal(fmt.Errorf("os.Readlink: linked path differs : '%s' != '%s'", target, link))
	}

	file := filepath.Join(link, "junction.go")
	fd, err := os.Open(file)
	if err != nil {
		t.Fatal(fmt.Errorf("mount may be fake because '%s' can not be open", file))
	}
	t.Log("file checked")
	defer fd.Close()
}
