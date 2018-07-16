package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/yuki-toida/knowme/config"
	"github.com/yuki-toida/knowme/interface/handler"
	"github.com/yuki-toida/knowme/interface/middleware/session"
)

func init() {
	config.Initialize()
	// model.Initialize()
}

func main() {
	router := gin.Default()
	session.AddMiddleware(router)

	router.StaticFS("/static", http.Dir("static"))
	router.LoadHTMLFiles("index.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	router.GET("/initial", func(c *gin.Context) { handler.GETInitial(c) })
	router.POST("/signin", func(c *gin.Context) { handler.POSTSignIn(c) })
	router.DELETE("/signout", func(c *gin.Context) { handler.DELETESignOut(c) })
	users := router.Group("/users")
	{
		users.GET("/events", func(c *gin.Context) { handler.GETUserEvent(c) })
		users.GET("/search/:id", func(c *gin.Context) { handler.GETUserSearch(c) })
	}
	events := router.Group("/events")
	{
		events.POST("/", func(c *gin.Context) { handler.POSTEvent(c) })
		events.PUT("/", func(c *gin.Context) { handler.PUTEvent(c) })
	}

	router.Run(":" + config.Config.Server.Port)
}
