package docker

import (
	"fmt"
	"testing"

	"github.com/muka/mufaas/util"
	"github.com/rs/xid"
)

func TestBuild(t *testing.T) {
	_, imageID := createImage(t, "hello")
	removeImage(t, imageID)
}

func TestBuildWithType(t *testing.T) {

	dir := "type-node1"
	uniqid := xid.New().String()
	imageName := fmt.Sprintf("%s-%s", dir, uniqid)
	srcPath := fmt.Sprintf("../test/%s", dir)

	err := util.CreateTar(srcPath)
	if err != nil {
		t.Fatal(err)
	}

	info, err := ImageBuild(ImageBuildOptions{
		Name:      imageName,
		Archive:   srcPath + ".tar",
		Type:      "node",
		TypesPath: []string{"../templates"},
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

	removeImage(t, info.ID)

}
