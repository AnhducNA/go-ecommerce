package main

import (
	"fmt"

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

func main() {
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
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	fmt.Println("Server port::", config.Server.Port)
	for _, db := range config.Databases {
		fmt.Println("Server database url: ", db.URL)
	}
}
