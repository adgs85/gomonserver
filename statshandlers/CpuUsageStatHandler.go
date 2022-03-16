package statshandlers

import (
	"net/http"

	"github.com/adgs85/gomonserver/monserver"
	"github.com/davecgh/go-spew/spew"
)

const CpuStatsType = "cpu"

func getCpuRegisterHandler() monserver.RegisterHandleFunc {
	requestPath := monserver.StatsEndpointPattern + "cpu"
	return getRegisterHandler(CpuStatsType, requestPath, cpuHandler{})
}

type cpuHandler struct {
}

func (cpuHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	stat := monserver.UnmarshalBodyToStat(r)
	println(spew.Sdump(stat))
}
