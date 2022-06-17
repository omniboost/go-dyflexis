package dyflexis_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	"github.com/omniboost/go-dyflexis"
)

func TestInformationstreamPut(t *testing.T) {
	req := client.NewInformationstreamPutRequest()
	req.PathParams().Key = "reservations"
	req.PathParams().Date = dyflexis.Date{time.Now()}
	req.RequestBody().Day.Value = 12
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
