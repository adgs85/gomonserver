package statshandlers

import (
	"encoding/json"
	"net/http"

	"github.com/adgs85/gomonserver/monserver"
	"github.com/adgs85/gomonserver/statdatabase"
)

const SnapshotStatsType = "snapshot"

func getStatSnapshotRegisterHandler() monserver.RegisterHandleFunc {
	requestPath := monserver.StatsEndpointPattern + SnapshotStatsType
	return getRegisterHandler(CpuStatsType, requestPath, statSnapshotHander{})
}

type statSnapshotHander struct {
}

func (statSnapshotHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if http.MethodGet != r.Method {
		w.WriteHeader(405)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	json, err := json.Marshal(statdatabase.FindAllSnapshot())
	if err != nil {
		w.WriteHeader(500)
		panic(err)
	}
	_, err = w.Write(json)
	monserver.PanicOnError(err)
}
