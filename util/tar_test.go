package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestCreateArchiveNonExistantDir(t *testing.T) {

	err := CreateTar("/foo")
	if err == nil {
		t.Fatal("Expeceted to fail creation")
	}
}

func createTar() (string, string, error) {

	root, err := ioutil.TempDir("", "test")
	if err != nil {
		return root, "", err
	}

	log.Debugf("Created tmp dir %s", root)

	dir, err := ioutil.TempDir(root, "content")
	if err != nil {
		return root, "", err
	}

	dir2, err := ioutil.TempDir(dir, "subdir")
	if err != nil {
		return root, "", err
	}

	err = ioutil.WriteFile(filepath.Join(dir, "file.txt"), []byte{}, 0644)
	if err != nil {
		return root, "", err
	}

	err = ioutil.WriteFile(filepath.Join(dir2, "file2.txt"), []byte{}, 0644)
	if err != nil {
		return root, "", err
	}

	// ignore: /subdir/file.tmp
	err = ioutil.WriteFile(filepath.Join(dir2, "file.tmp"), []byte{}, 0644)
	if err != nil {
		return root, "", err
	}

	// ignore: /file.tmp
	err = ioutil.WriteFile(filepath.Join(dir, "file.tmp"), []byte{}, 0644)
	if err != nil {
		return root, "", err
	}

	err = ioutil.WriteFile(filepath.Join(dir, ".dockerignore"), []byte("file.tmp\n"), 0644)
	if err != nil {
		return root, "", err
	}

	err = CreateTar(dir)
	if err != nil {
		return root, "", err
	}
	_, err = os.Stat(dir + ".tar")
	if err != nil {
		return root, "", err
	}

	return root, dir + ".tar", err
}

func TestCreateTar(t *testing.T) {
	root, _, err := createTar()
	defer os.Remove(root)
	if err != nil {
		t.Fatal(err)
	}
}

func TestExtractTar(t *testing.T) {
	root, tar, err := createTar()
	defer os.Remove(root)
	if err != nil {
		t.Fatal(err)
	}

	dir, err := ioutil.TempDir("", "test")
	defer os.Remove(dir)
	if err != nil {
		t.Fatal(err)
	}

	err = ExtractTar(tar, dir)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(filepath.Join(dir, "subdir", "file.tmp")); err == nil {
		t.Fatal("file.tmp should have been ignored")
	}

	if _, err := os.Stat(filepath.Join(dir, "file.tmp")); err == nil {
		t.Fatal("file.tmp should have been ignored")
	}

}
