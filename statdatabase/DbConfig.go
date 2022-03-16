package statdatabase

import (
	"fmt"

	"github.com/adgs85/gomonmarshalling/monmarshalling/envconfig"
)

type dbCofig struct {
	Host     string `mapstructure:"db_host"`
	Port     int    `mapstructure:"db_port"`
	User     string `mapstructure:"db_user"`
	Password string `mapstructure:"db_password"`
	Name     string `mapstructure:"db_name"`
}

var _, psqlConnectionStr = initCfg()

func initCfg() (*dbCofig, string) {
	c := &dbCofig{}
	envconfig.GetViperConfig().Unmarshal(c)
	//println(spew.Sdump(envconfig.GetViperConfig()))
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Name)
	return c, psqlconn
}
