package docker

import (
	"os"
	"testing"

	"github.com/muka/mufaas/util"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
)

func TestMain(m *testing.M) {
	log.SetLevel(log.DebugLevel)
	os.Exit(m.Run())
}

func TestFailListFilter(t *testing.T) {
	_, err := ImageList([]string{"foo,bar"})
	if err == nil {
		t.Fatal("Filter error expected")
	}
}

func TestBuild(t *testing.T) {
	imageName := "mufaas/hello-" + xid.New().String()
	imageID := doBuild(t, "../test/hello", imageName)
	err := ImageRemove(imageID, true)
	if err != nil {
		t.Fatal(err)
	}
}

// Build builds a docker image from the image directory
func doBuild(t *testing.T, srcPath, imageName string) string {

	var err error

	//HACK: docker complain about the symlink in node_modules/.bin
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
