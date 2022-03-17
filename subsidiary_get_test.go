package timesheet_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestSubsidiaryGet(t *testing.T) {
	req := client.NewSubsidiaryGetRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
