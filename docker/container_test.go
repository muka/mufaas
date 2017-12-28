package docker

import (
	"testing"
)

func TestCreate(t *testing.T) {
	info := createContainer(t, "hello")
	removeContainer(t, info.ID)
	removeImage(t, info.ImageID)
}
