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

// ConnectDB func
func ConnectDB() *gorm.DB {
	connectionString := "root:zaqroot@tcp(" + Config.Db.Host + ":" + Config.Db.Port + ")/" + Config.Db.Name + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	db.LogMode(true)
	return db
}
