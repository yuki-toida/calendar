package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/yuki-toida/knowme/config"
	"github.com/yuki-toida/knowme/domain/model"
	"github.com/yuki-toida/knowme/infrastructure/repository/event"
	"github.com/yuki-toida/knowme/infrastructure/repository/user"
	"github.com/yuki-toida/knowme/interface/handler"
	"github.com/yuki-toida/knowme/interface/middleware/session"
	"github.com/yuki-toida/knowme/interface/registry"
)

func init() {
	config.Initialize()
}

func main() {
	connectionString := "root:zaqroot@tcp(" + config.Config.Db.Host + ":" + config.Config.Db.Port + ")/" + config.Config.Db.Name + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	db.LogMode(true)
	db.AutoMigrate(&model.User{}, &model.Event{})
	model.DB = db

	registry := registry.New(user.New(db), event.New(db))

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

	router.GET("/initial", func(c *gin.Context) { handler.GETInitial(c, registry) })
	router.POST("/signin", func(c *gin.Context) { handler.POSTSignIn(c, registry) })
	router.DELETE("/signout", func(c *gin.Context) { handler.DELETESignOut(c, registry) })
	router.GET("/search/:id", func(c *gin.Context) { handler.GETSearch(c, registry) })

	events := router.Group("/events")
	{
		events.GET("/", func(c *gin.Context) { handler.GETEvent(c, registry) })
		events.POST("/", func(c *gin.Context) { handler.POSTEvent(c, registry) })
		events.PUT("/", func(c *gin.Context) { handler.PUTEvent(c, registry) })
	}

	router.Run(":" + config.Config.Server.Port)
}
