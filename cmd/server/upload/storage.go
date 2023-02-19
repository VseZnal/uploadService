package upload

import (
	"io/ioutil"
)

type Manager interface {
	Store(file *File) error
}

var _ Manager = &Storage{}

type Storage struct {
	dir string
}

func New(dir string) Storage {
	return Storage{
		dir: dir,
	}
}

func (s Storage) Store(file *File) error {
	if err := ioutil.WriteFile(s.dir+file.Name, file.Buffer.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}
