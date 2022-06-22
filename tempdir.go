package tugo

import (
	"os"
)

type tempDir struct {
	path string
}

// Path is the getter of the temporary directory path.
func (d *tempDir) Path() string {
	return d.path
}

// Remove removes the temporary directory and any children it contains.
func (d *tempDir) Remove() error {
	return os.RemoveAll(d.path)
}

type ITempDir interface {
	// Path returns the temporary directory path.
	Path() string

	// Remove is a wrapper of os RemoveAll method. It removes the
	// temporary directory and any children it contains.
	Remove() error
}

// TempDir is a wrapper of os MkdirTemp method from which it inherits
// the same arguments: https://pkg.go.dev/os?utm_source=gopls#MkdirTemp.
// It returns ITempDir, an interface withe the directory path getter and
// the method to remove it. It is caller's reponsability to call Remove
// when it is no longer needed.
func TempDir(dir string, pattern string) (ITempDir, error) {
	var err error
	d := new(tempDir)
	d.path, err = os.MkdirTemp(dir, pattern)
	if err != nil {
		return nil, err
	}
	return d, nil
}
