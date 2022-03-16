package statshandlers

import (
	"log"
	"net/http"

	"github.com/adgs85/gomonserver/monserver"
	"github.com/davecgh/go-spew/spew"
)

var requestPath = monserver.StatsEndpointPattern + "cpu"

func getRegisterHandler() monserver.RegisterHandleFunc {
	return func(serverMux *http.ServeMux) {
		log.Println("Registering Disk stats handler on", requestPath)
		serverMux.Handle(requestPath, cpuHandler{})
	}
}

type cpuHandler struct {
}

func (cpuHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	stat := monserver.UnmarshalBodyToStat(r)
	println(spew.Sdump(stat))
}
