package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

const domain = "@candee.co.jp"

// Config struct
var Config struct {
	Env    string
	Domain string
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
	Config.Domain = domain
}
