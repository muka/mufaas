package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestCreateArchiveNonExistantDir(t *testing.T) {

	err := CreateArchive("/foo")
	if err == nil {
		t.Fatal("Expeceted to fail creation")
	}
}

func TestCreateArchive(t *testing.T) {

	root, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(root)

	log.Debugf("Created tmp dir %s", root)

	dir, err := ioutil.TempDir(root, "content")
	if err != nil {
		t.Fatal(err)
	}

	err = ioutil.WriteFile(filepath.Join(dir, "file.txt"), []byte{}, 0x700)
	if err != nil {
		t.Fatal(err)
	}

	err = CreateArchive(dir)
	if err != nil {
		t.Fatal(err)
	}

	_, err = os.Stat(dir + ".tar")
	if err != nil {
		t.Fatal(err)
	}

}
