package tugo

import (
	"os"
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
