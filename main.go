package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/yuki-toida/knowme/config"
	"github.com/yuki-toida/knowme/controller"
	"github.com/yuki-toida/knowme/model"
)

func init() {
	config.Initialize()
	model.Migrate()
}

func main() {
	router := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("_knowme", store))

	router.LoadHTMLFiles("index.html")

	if config.Config.Env == "local" {
		router.StaticFS("/static", http.Dir("static"))
	}

	router.GET("/healthz", controller.Healthz)
	router.GET("/", controller.Index)
	router.GET("/initial", controller.Initial)
	router.POST("/signin", controller.SignIn)
	router.POST("/signout", controller.SignOut)
	router.POST("/events", controller.AddEvent)
	router.PUT("/events", controller.DeleteEvent)

	router.Run(":" + config.Config.Server.Port)
}
