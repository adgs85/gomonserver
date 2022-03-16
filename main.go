package main

import (
	"github.com/adgs85/gomonserver/monserver"
	"github.com/adgs85/gomonserver/statshandlers"
)

func main() {
	monserver.StartHttpServer(statshandlers.NewHandlerList())
}
