package monserver

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/adgs85/gomonmarshalling/monmarshalling"
)

func UnmarshalBodyToStat(r *http.Request) *monmarshalling.Stat {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error while reading request bodu", err)
		return nil
	} else {
		return UnmarshalToStat(body)
	}
}

func UnmarshalToStat(body []byte) *monmarshalling.Stat {
	stat := monmarshalling.Stat{}
	json.Unmarshal(body, &stat)
	return &stat
}
