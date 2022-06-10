package file

import (
	"io/fs"
	"os"
	"path"
)

type File struct {
	fs.DirEntry

	Root string
	Path string
}

func NewFile(entry fs.DirEntry, root, p string) *File {
	return &File{
		DirEntry: entry,

		Root: root,
		Path: path.Join(root, p),
	}
}

func (f File) Ext() string {
	return path.Ext(f.Name())
}

type WalkFunc func(*File) error

func Walk(root string, fn WalkFunc) error {
	return fs.WalkDir(
		os.DirFS(root), ".",
		func(path string, d fs.DirEntry, err error) error {
			return fn(NewFile(d, root, path))
		},
	)
}
