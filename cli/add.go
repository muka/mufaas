package cli

import (
	"github.com/jhoonb/archivex"
)

func CreateArchive(dir string) error {

	tar := new(archivex.TarFile)
	err := tar.Create(dir)
	if err != nil {
		return err
	}
	err = tar.AddAll(dir, false)
	if err != nil {
		return err
	}
	err = tar.Close()
	if err != nil {
		return err
	}

	return nil
}
