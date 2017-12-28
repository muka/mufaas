package docker

import (
	"testing"
)

func TestFailListFilter(t *testing.T) {
	_, err := ImageList([]string{"foo,bar"})
	if err == nil {
		t.Fatal("Filter error expected")
	}
}

func TestBuild(t *testing.T) {
	_, imageID := createImage(t, "hello")
	removeImage(t, imageID)
}
