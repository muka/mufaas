package template

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func CreateFunction(sourcePath string, funcType string, typesPath []string) (string, error) {

	if funcType == "" {
		return "", errors.New("A type or Dockerfile has not been provided")
	}
	if len(typesPath) == 0 {
		return "", errors.New("Types path is empty")
	}

	typePath := getBasePath(funcType, typesPath)
	if typePath == "" {
		return "", fmt.Errorf("Type %s not found", funcType)
	}

	dst, err := ioutil.TempDir("", "mufaas_build")
	if err != nil {
		return "", err
	}

	// copy template first
	err = copy(typePath, dst)
	if err != nil {
		return "", err
	}

	// copy user function file then
	err = copy(sourcePath, dst)
	if err != nil {
		return "", err
	}

	return dst, nil
}

func getBasePath(funcType string, typesPath []string) string {
	for _, p := range typesPath {
		fpath := filepath.Join(p, funcType, "Dockerfile")
		log.Debugf("Lookup for %s", fpath)
		if _, err := os.Stat(fpath); err == nil {
			return filepath.Join(p, funcType)
		}
	}
	return ""
}

func copy(src, dst string) error {
	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return nil
	}
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())
		if entry.IsDir() {
			err = CopyDir(srcPath, dstPath)
			if err != nil {
				return nil
			}
		} else {
			// Skip symlinks.
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}
			err = CopyFile(srcPath, dstPath)
			if err != nil {
				return nil
			}
		}
	}
	return nil
}
