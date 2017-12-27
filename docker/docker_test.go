package docker

import (
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

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

	f := []string{"label=mufaas=1"}
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
