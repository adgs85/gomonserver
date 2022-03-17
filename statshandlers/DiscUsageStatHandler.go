package statshandlers

import (
	"github.com/adgs85/gomonserver/monserver"
)

const DiskStatsType = "disk"

func getDiskRegisterHandler() monserver.RegisterHandleFunc {
	requestPath := monserver.StatsEndpointPattern + DiskStatsType
	return getRegisterHandler(CpuStatsType, requestPath, statHandler{})
}
