package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/yuki-toida/knowme/config"
	"github.com/yuki-toida/knowme/infrastructure/repository"
	"github.com/yuki-toida/knowme/interface/controller"
	"github.com/yuki-toida/knowme/interface/middleware"
	"github.com/yuki-toida/knowme/registry"
)

func main() {
	config.Init()

	connectionString := "root:zaqroot@tcp(" + config.Config.Db.Host + ":" + config.Config.Db.Port + ")/" + config.Config.Db.Name + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	db.LogMode(true)
	// db.AutoMigrate(&user.User{}, &event.Event{})
	// defer db.Close()

	registry := registry.Registry{
		UserRepository: user.New(db),
	}

	router := gin.Default()

	router.Use(middleware.AddRegistry(registry))
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("_knowme", store))
	router.Use(middleware.AddSession())

	router.StaticFS("/static", http.Dir("static"))
	router.LoadHTMLFiles("index.html")

	router.GET("/healthz", controller.HomeHealthz)
	router.GET("/", controller.HomeIndex)
	router.GET("/initial", controller.HomeInitial)
	router.POST("/signin", controller.HomeSignIn)
	router.DELETE("/signout", controller.HomeSignOut)
	users := router.Group("/users")
	{
		users.GET("/events", controller.UserEvents)
		users.GET("/search/:id", controller.User)
	}
	events := router.Group("/events")
	{
		events.POST("/", controller.EventAdd)
		events.PUT("/", controller.EventDelete)
	}

	router.Run(":" + config.Config.Server.Port)
}
