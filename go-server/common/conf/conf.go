package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

var Viper = viper.New()

func InitConfig() {
	var err error
	Viper.SetConfigName("app")
	Viper.SetConfigType("toml")
	Viper.AddConfigPath("go_server/conf/dev")
	err = Viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}