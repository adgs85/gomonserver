package monserver

import (
	"log"
	"net/http"
)

func StartHttpServer(registerFunctionList []RegisterHandleFunc) {
	cfg := GlobalCfg()
	bindAddress := cfg.BindIp + ":" + cfg.ServerPort
	log.Println("Starting http Server on", bindAddress)

	statsServerMux := http.NewServeMux()
	for _, f := range registerFunctionList {
		f(statsServerMux)
	}
	statsServerMux.HandleFunc(HeartBeatEndpointPath, heartbeatHandler)

	http.ListenAndServe(bindAddress, statsServerMux)

}

func heartbeatHandler(w http.ResponseWriter, r *http.Request) {
	UnmarshalBodyToStat(r)
	//println(spew.Sdump(stat))
}
