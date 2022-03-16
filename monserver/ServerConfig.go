package monserver

import (
	"net/http"

	"github.com/adgs85/gomonmarshalling/monmarshalling/envconfig"
)

var globalCfg = initCfg()

const StatsEndpointPattern = "/stats/"

const HeartBeatEndpointPath = "/heartbeat"

type RegisterHandleFunc = func(serverMux *http.ServeMux)

type serverConfig struct {
	BindIp     string `mapstructure:"server_bind_ip"`
	ServerPort string `mapstructure:"server_port"`
}

func GlobalCfg() *serverConfig {
	return globalCfg
}

func initCfg() *serverConfig {
	cfg := new(serverConfig)
	envconfig.GetViperConfig().Unmarshal(cfg)
	return cfg
}
