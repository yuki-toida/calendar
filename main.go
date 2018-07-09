package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/yuki-toida/knowme/config"
	"github.com/yuki-toida/knowme/controller"
	"github.com/yuki-toida/knowme/model"
)

func sessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.URL.Path, "/static") {
			session := sessions.Default(c)
			if id := session.Get(config.SessionName); id != nil {
				db := config.ConnectDB()
				user := model.GetUser(db, id.(string))
				c.Set(config.SessionName, user)
			}
		}
	}
}

func init() {
	config.Initialize()
	model.Migrate()
}

func main() {
	router := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))

	router.Use(sessions.Sessions("_knowme", store))
	router.Use(sessionMiddleware())

	router.LoadHTMLFiles("index.html")
	if config.Config.Env == "local" {
		router.StaticFS("/static", http.Dir("static"))
	}

	router.GET("/healthz", controller.HomeHealthz)
	router.GET("/", controller.HomeIndex)
	router.GET("/initial", controller.HomeInitial)
	router.POST("/signin", controller.HomeSignIn)
	router.DELETE("/signout", controller.HomeSignOut)

	event := router.Group("/events")
	{
		event.POST("/", controller.EventAdd)
		event.PUT("/", controller.EventDelete)
	}

	router.Run(":" + config.Config.Server.Port)
}
