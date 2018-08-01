package main

import (
	"fmt"
	"net/http"
	"time"

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
	config.Init()
}

func main() {
	connectionString := "root:zaqroot@tcp(" + config.Config.Db.Host + ":" + config.Config.Db.Port + ")/" + config.Config.Db.Name + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	db.LogMode(true)
	db.AutoMigrate(&model.User{}, &model.Event{})

	registry := registry.NewRegistry(user.NewRepository(db), event.NewRepository(db))
	handler := handler.NewHandler(registry)

	router := gin.Default()
	session.AddMiddleware(router)
	router.StaticFS("/static", http.Dir("static"))
	router.LoadHTMLGlob("interface/template/*")

	router.GET("/", func(c *gin.Context) {
		now := time.Now()
		revision := fmt.Sprintf("%d%d%d%d%d", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute())
		c.HTML(http.StatusOK, "index.html", gin.H{
			"revision": revision,
		})
	})
	router.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	router.GET("/initial", func(c *gin.Context) { handler.Initial(c) })
	router.POST("/signin", func(c *gin.Context) { handler.SignIn(c) })
	router.DELETE("/signout", func(c *gin.Context) { handler.SignOut(c) })
	router.GET("/search/:id", func(c *gin.Context) { handler.Search(c) })
	router.GET("/images", func(c *gin.Context) { handler.Images(c) })
	router.POST("/upload", func(c *gin.Context) { handler.Upload(c) })

	events := router.Group("/events")
	{
		events.GET("/", func(c *gin.Context) { handler.GetEvent(c) })
		events.POST("/", func(c *gin.Context) { handler.CreateEvent(c) })
		events.PUT("/", func(c *gin.Context) { handler.DeleteEvent(c) })
	}

	router.Run(":" + config.Config.Server.Port)
}
