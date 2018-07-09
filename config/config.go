package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
)

// SessionName const
const SessionName = "ID"

// Config struct
var Config struct {
	Env    string
	Server struct {
		Host       string `toml:"host"`
		Port       string `toml:"port"`
		StaticPath string `toml:"static-path"`
		BucketName string `toml:"bucket-name"`
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

// ConnectDB func
func ConnectDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:zaqroot@tcp(127.0.0.1:7306)/knowme?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	db.LogMode(true)
	return db
}
