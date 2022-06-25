package tugo

import (
	"os"
	"path/filepath"
	"testing"
)

func TestTempDir(t *testing.T) {
	d, err := TempDir(".", "testing")
	if err != nil {
		t.Fatal("failed to create temporary directory:", err)
	}

	if err := d.Remove(); err != nil {
		t.Fatal("failed to remove temporary directory:", err)
	}

	if _, err := os.Open(d.Path()); err == nil {
		t.Fatal("temporary directory has not been removed")
	}
}

func TestTempDirIgnore(t *testing.T) {
	d, err := TempDir(".", "testing-ignore", IgnoreDir())
	if err != nil {
		t.Fatal("failed to create temporary directory:", err)
	}

	if f, err := os.Open(filepath.Join(d.Path(), ".gitignore")); err != nil {
		t.Error(".gitignore file not found in the temporary directory")
	} else {
		_ = f.Close()
	}

	if err := d.Remove(); err != nil {
		t.Fatal("failed to remove temporary directory:", err)
	}

	if _, err := os.Open(d.Path()); err == nil {
		t.Fatal("temporary directory has not been removed")
	}
}
