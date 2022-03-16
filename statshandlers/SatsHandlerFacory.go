package statshandlers

import (
	"log"
	"net/http"

	"github.com/adgs85/gomonserver/monserver"
)

func NewHandlerList() []monserver.RegisterHandleFunc {
	return []monserver.RegisterHandleFunc{
		getCpuRegisterHandler(),
	}
}

func getRegisterHandler(statType string, requestPath string, handler http.Handler) monserver.RegisterHandleFunc {
	return func(serverMux *http.ServeMux) {
		log.Println("Registering", statType, "stats handler on", requestPath)
		serverMux.Handle(requestPath, handler)
	}
}
