package docker

import (
	"fmt"
	"os"
	"testing"

	"github.com/muka/mufaas/util"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
)

type containerInfo struct {
	Name      string
	ID        string
	ImageName string
	ImageID   string
}

func TestMain(m *testing.M) {

	var v bool
	for _, arg := range os.Args {
		if arg == "-test.v=true" {
			v = true
			break
		}
	}

	if v {
		log.SetLevel(log.DebugLevel)
	}

	f := []string{"label=" + DefaultLabel + "=1"}
	list, err := ImageList(f)
	if err != nil {
		panic(err)
	}

	for _, image := range list {
		err := ImageRemove(image.ID, true)
		if err != nil {
			log.Warnf("Fail remove: %s", err.Error())
		}
	}

	log.Debugf("Removed previous images")

	os.Exit(m.Run())
}

func createImage(t *testing.T, dir string) (string, string) {

	uniqid := xid.New().String()
	imageName := fmt.Sprintf("%s-%s", dir, uniqid)
	imageID := doBuild(t, fmt.Sprintf("../test/%s", dir), imageName)

	log.Debugf("Created image %s %s", imageName, imageID)
	return imageName, imageID
}

func removeImage(t *testing.T, imageID string) {
	err := ImageRemove(imageID, true)
	if err != nil {
		t.Fatal(err)
	}
}

func removeContainer(t *testing.T, containerID string) {
	err := Remove(containerID, true)
	if err != nil {
		t.Fatal(err)
	}
}

func createContainer(t *testing.T, dir string) containerInfo {

	imageName, imageID := createImage(t, dir)
	opts := CreateOptions{
		Name:  imageName,
		Image: imageName,
	}

	info, err := Create(&opts)
	if err != nil {
		t.Fatalf("Create failed: %s", err.Error())
	}

	return containerInfo{
		Name:      info.Name,
		ID:        info.ID,
		ImageName: imageName,
		ImageID:   imageID,
	}
}

// Build builds a docker image from the image directory
func doBuild(t *testing.T, srcPath, imageName string) string {

	var err error

	err = util.CreateTar(srcPath)
	if err != nil {
		t.Fatal(err)
	}

	info, err := ImageBuild(ImageBuildOptions{
		Name:    imageName,
		Archive: srcPath + ".tar",
	})
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
