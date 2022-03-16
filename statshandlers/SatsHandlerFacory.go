package statshandlers

import "github.com/adgs85/gomonserver/monserver"

func NewHandlerList() []monserver.RegisterHandleFunc {
	return []monserver.RegisterHandleFunc{
		getRegisterHandler(),
	}
}
