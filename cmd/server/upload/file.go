package upload

import "bytes"

type File struct {
	Name   string
	Buffer *bytes.Buffer
}

func NewFile(name string) *File {
	return &File{
		Name:   name,
		Buffer: &bytes.Buffer{},
	}
}

func (f *File) Write(chunk []byte) error {
	_, err := f.Buffer.Write(chunk)

	return err
}
