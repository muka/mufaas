package util

import (
	"os"
	"path/filepath"

	"github.com/mholt/archiver"
	"github.com/monochromegane/go-gitignore"
)

func CreateTar(dir string) error {

	list := []string{dir}

	ignfile := filepath.Join(dir, ".dockerignore")
	if _, err := os.Stat(ignfile); err == nil {

		gitignore, err := gitignore.NewGitIgnore(ignfile, dir)
		if err != nil {
			return err
		}

		list := []string{}
		err = filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {

			if !gitignore.Match(path, f.IsDir()) {
				list = append(list, path)
			}

			return nil
		})

		if err != nil {
			return err
		}

	}

	return archiver.Tar.Make(dir+".tar", list)
}

func ExtractTar(archive, dst string) error {
	return archiver.Tar.Open(archive, dst)
}
