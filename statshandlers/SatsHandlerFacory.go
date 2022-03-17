package statshandlers

import (
	"log"
	"net/http"

	"github.com/adgs85/gomonserver/monserver"
	"github.com/adgs85/gomonserver/statdatabase"
)

func NewHandlerList() []monserver.RegisterHandleFunc {
	return []monserver.RegisterHandleFunc{
		getCpuRegisterHandler(),
		getDiskRegisterHandler(),
	}
}

type statHandler struct {
}

func (statHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	stat := monserver.UnmarshalBodyToStat(r)
	statdatabase.InsertStat(stat)
}

func getRegisterHandler(statType string, requestPath string, handler http.Handler) monserver.RegisterHandleFunc {
	return func(serverMux *http.ServeMux) {
		log.Println("Registering", statType, "stats handler on", requestPath)
		serverMux.Handle(requestPath, handler)
	}
}
