package statshandlers

import (
	"github.com/adgs85/gomonserver/monserver"
)

const CpuStatsType = "cpu"

func postCpuRegisterHandler() monserver.RegisterHandleFunc {
	requestPath := monserver.StatsEndpointPattern + CpuStatsType
	return getRegisterHandler(CpuStatsType, requestPath, statHandler{})
}
