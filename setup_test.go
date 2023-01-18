package dyflexis_test

import (
	"os"
	"testing"

	"github.com/omniboost/go-dyflexis"
)

var (
	client *dyflexis.Client
)

func TestMain(m *testing.M) {
	debug := os.Getenv("DEBUG")
	systemName := os.Getenv("SYSTEM_NAME")
	api2Token := os.Getenv("API2_TOKEN")

	client = dyflexis.NewClient(nil)
	if debug != "" {
		client.SetDebug(true)
	}

	client.SetSystemName(systemName)
	client.SetAPI2Token(api2Token)

	client.SetDisallowUnknownFields(true)
	m.Run()
}
