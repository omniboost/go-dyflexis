package dyflexis_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestEmployeesPost(t *testing.T) {
	req := client.NewEmployeesPostRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
