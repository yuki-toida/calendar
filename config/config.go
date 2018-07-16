package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

// SessionName const
const SessionName = "ID"

// Config struct
var Config struct {
	Env    string
	Server struct {
		Host string `toml:"host"`
		Port string `toml:"port"`
	}
	Db struct {
		Host string `toml:"host"`
		Port string `toml:"port"`
		Name string `toml:"name"`
	}
}

// Initialize func
func Initialize() {
	env := os.Getenv("ENV")
	_, err := toml.DecodeFile("config/"+env+".toml", &Config)
	if err != nil {
		panic(err)
	}
	Config.Env = env
	fmt.Printf("[CONFIG] : %+v\n", Config)
}
