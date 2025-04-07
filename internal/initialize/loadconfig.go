package initialize

import (
	"fmt"

	"github.com/AnhducNA/go-ecommerce/global"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		URL string `mapstructure:"url"`
	} `mapstructure:"database"`
}

func LoadConfig() {
	viper := viper.New()
	viper.SetConfigName("local") // name of config file
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config") //path to config
	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	// read server configuration
	fmt.Println("Server port::", viper.GetInt("server.port"))
	fmt.Println("Server database url: ", viper.GetString("database.url"))
	fmt.Println("Security jwt key: ", viper.GetString("security.jwt.key"))

	// configuration with struct
	err = viper.Unmarshal(&global.Config)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
}
