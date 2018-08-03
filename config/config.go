package config

import (
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

const domain = "@candee.co.jp"

// Config struct
var Config struct {
	Env    string
	Domain string
	Server struct {
		Host       string `toml:"host"`
		Port       string `toml:"port"`
		Bucket     string `toml:"bucket"`
		StorageURL string `toml:"storage-url"`
	}
	Db struct {
		Host string `toml:"host"`
		Port string `toml:"port"`
		Name string `toml:"name"`
	}
}

// JST *time.Location
var JST *time.Location

// Now time.Time
var Now time.Time

// Init func
func Init() {
	env := os.Getenv("ENV")
	_, err := toml.DecodeFile("config/"+env+".toml", &Config)
	if err != nil {
		panic(err)
	}
	Config.Env = env
	Config.Domain = domain

	JST = time.FixedZone("JST", 9*60*60)
	Now = time.Now().In(JST)
}
