package tugo

import (
	"os"
	"path/filepath"
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

// TempDirOpt is a TempDir functional option. Its argument
// is the temporary directory path.
type TempDirOpt func(string) error

// TempDir is a wrapper of os MkdirTemp method from which it inherits
// the same arguments: https://pkg.go.dev/os?utm_source=gopls#MkdirTemp.
// It returns ITempDir, an interface withe the directory path getter and
// the method to remove it. It is caller's reponsability to call Remove
// when it is no longer needed.
func TempDir(dir string, pattern string, opts ...TempDirOpt) (ITempDir, error) {
	var err error
	d := new(tempDir)
	d.path, err = os.MkdirTemp(dir, pattern)
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		err = opt(d.path)
		if err != nil {
			_ = d.Remove()
			return nil, err
		}
	}

	return d, nil
}

// IgnoreDir is an option for TempDir that writes a .gitignore file
// with content "*" to ignore all files in the temporary directory.
func IgnoreDir() TempDirOpt {
	return func(d string) error {
		data := []byte("*")
		return os.WriteFile(filepath.Join(d, ".gitignore"), data, 0766)
	}
}
