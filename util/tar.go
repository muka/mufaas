package util

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mholt/archiver"
	"github.com/monochromegane/go-gitignore"
	log "github.com/sirupsen/logrus"
)

func CreateTar(dir string) error {

	log.Debugf("Create archive for %s", dir)

	ignfile := filepath.Join(dir, ".dockerignore")
	var ignore gitignore.IgnoreMatcher
	if _, err := os.Stat(ignfile); err == nil {
		log.Debugf("Found %s", ignfile)
		ignore, err = gitignore.NewGitIgnore(ignfile, dir)
		if err != nil {
			return err
		}
	}

	list, err := readDir(dir, ignore)
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

func readDir(basePath string, ignore gitignore.IgnoreMatcher) ([]string, error) {
	list := []string{}
	entries, err := ioutil.ReadDir(basePath)
	if err != nil {
		return list, err
	}
	for _, entry := range entries {

		fullpath := filepath.Join(basePath, entry.Name())
		if ignore != nil {
			if !ignore.Match(fullpath, entry.IsDir()) {
				list = append(list, fullpath)
			} else {
				log.Debugf("Skip %s", fullpath)
				continue
			}
		} else {
			list = append(list, fullpath)
		}

		if entry.IsDir() {
			list2, err := readDir(filepath.Join(basePath, entry.Name()), ignore)
			if err != nil {
				return list, err
			}
			list = append(list, list2...)
		}

	}

	return list, nil
}
