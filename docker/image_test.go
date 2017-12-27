package docker

import (
	"os"
	"testing"

	"github.com/muka/mufaas/util"
	log "github.com/sirupsen/logrus"
)

func TestMain(m *testing.M) {
	log.SetLevel(log.DebugLevel)
	os.Exit(m.Run())
}

func TestBuild(t *testing.T) {
	doBuild(t, "../test/hello", "mufaas/hello")
}

// Build builds a docker image from the image directory
func doBuild(t *testing.T, srcPath, imageName string) string {

	var err error

	//HACK: for some reason docker complain about the symlink in node_modules/.bin
	dotbin := srcPath + "/node_modules/.bin"
	if _, serr := os.Stat(dotbin); !os.IsNotExist(serr) {
		err = os.RemoveAll(dotbin)
		if err != nil {
			t.Fatal(err)
		}
	}

	err = util.CreateArchive(srcPath)
	if err != nil {
		t.Fatal(err)
	}

	info, err := ImageBuild(imageName, srcPath+".tar")
	if err != nil {
		t.Fatal(err)
	}

	list, err := ImageList([]string{"reference=" + imageName})

	if err != nil {
		t.Fatal(err)
	}

	if len(list) == 0 {
		t.Fatal("Found 0 images")
	}

	return info.ID
}
