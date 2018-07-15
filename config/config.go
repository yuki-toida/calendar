package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

const emailDomain = "@candee.co.jp"

// Config struct
var Config struct {
	Env         string
	EmailDomain string
	Server      struct {
		Host string `toml:"host"`
		Port string `toml:"port"`
	}
	Db struct {
		Host string `toml:"host"`
		Port string `toml:"port"`
		Name string `toml:"name"`
	}
}

// Init func
func Init() {
	env := os.Getenv("ENV")
	_, err := toml.DecodeFile("config/"+env+".toml", &Config)
	if err != nil {
		panic(err)
	}
	Config.Env = env
	Config.EmailDomain = emailDomain
	fmt.Printf("[CONFIG] : %+v\n", Config)
}
