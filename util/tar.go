package util

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/mholt/archiver"
	"github.com/monochromegane/go-gitignore"
	log "github.com/sirupsen/logrus"
)

func CreateTar(dir string) error {

	log.Debugf("Create archive for %s", dir)

	list := []string{}

	ignfile := filepath.Join(dir, ".dockerignore")

	var ignore gitignore.IgnoreMatcher
	if _, err := os.Stat(ignfile); err == nil {
		log.Debugf("Found %s", ignfile)
		ignore, err = gitignore.NewGitIgnore(ignfile)
		if err != nil {
			return err
		}
	}

	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {

		if path == dir {
			return nil
		}

		if ignore != nil {
			if !ignore.Match(path, f.IsDir()) {
				// log.Debugf("Add %s", path)
				list = append(list, path)
			}
		} else {
			// log.Debugf("Add %s", path)
			list = append(list, path)
		}

		return nil
	})

	if err != nil {
		return err
	}

	if len(list) == 0 {
		return errors.New("Cannot create an empty archive")
	}

	return archiver.Tar.Make(dir+".tar", list)
}

func ExtractTar(archive, dst string) error {
	return archiver.Tar.Open(archive, dst)
}
