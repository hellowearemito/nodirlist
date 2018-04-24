package nodirlist

import (
	"net/http"
	"os"
)

// FS is a http.FileSystem that doesn't allow
// access to directories.
type nolistfs struct {
	fs http.FileSystem
}

// Wrap wraps a http.FileSystem and disallows access to directories.
func Wrap(fs http.FileSystem) http.FileSystem {
	return &nolistfs{
		fs: fs,
	}
}

func (fs *nolistfs) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if fi.IsDir() {
		return nil, os.ErrPermission
	}
	return f, nil
}
