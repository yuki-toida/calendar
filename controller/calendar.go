package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuki-toida/knowme/config"
	"github.com/yuki-toida/knowme/model"
)

// CalendarHealthz func
func CalendarHealthz(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

// Calendar func
func Calendar(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"staticUrl": config.Config.Server.StaticURL})
}

// CalendarInit func
func CalendarInit(c *gin.Context) {
	var params struct {
		Email string `json:"email"`
		Name  string `json:"name"`
		Photo string `json:"photo"`
	}
	if err := c.ShouldBindJSON(&params); err == nil {
		user, events := model.GetBase(params.Email, params.Name, params.Photo)
		c.JSON(http.StatusOK, gin.H{
			"staticUrl": config.Config.Server.StaticURL,
			"user":      user,
			"events":    events,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// CalendarAdd func
func CalendarAdd(c *gin.Context) {
	var params struct {
		UserID string    `json:"userId"`
		Date   time.Time `json:"date"`
	}
	if err := c.ShouldBindJSON(&params); err == nil {
		event := model.AddEvent(params.UserID, params.Date.In(time.Local))
		c.JSON(http.StatusOK, gin.H{"event": event})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// CalendarDelete func
func CalendarDelete(c *gin.Context) {
	var params struct {
		EventID string `json:"id"`
	}
	if err := c.ShouldBindJSON(&params); err == nil {
		model.DeleteEvent(params.EventID)
		c.JSON(http.StatusOK, gin.H{"ok": true})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
