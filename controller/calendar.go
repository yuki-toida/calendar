package controller

import (
	"net/http"
	"strings"
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
	if err := c.ShouldBindJSON(&params); err != nil {
		handleError(c, err)
	} else {
		var user model.User
		if params.Email != "" && strings.Contains(params.Email, model.EmailDomain) {
			user = model.InitUser(params.Email, params.Name, params.Photo)
		}
		c.JSON(http.StatusOK, gin.H{
			"staticUrl":   config.Config.Server.StaticURL,
			"emailDomain": model.EmailDomain,
			"events":      model.GetEvents(),
			"user":        user,
		})
	}
}

// CalendarAdd func
func CalendarAdd(c *gin.Context) {
	var params struct {
		UserID string    `json:"userId"`
		Date   time.Time `json:"date"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		handleError(c, err)
	} else {
		event, err := model.AddEvent(params.UserID, params.Date.In(time.Local))
		if err != nil {
			handleError(c, err)
		} else {
			c.JSON(http.StatusOK, gin.H{"event": event})
		}
	}
}

// CalendarDelete func
func CalendarDelete(c *gin.Context) {
	var params struct {
		EventID string `json:"id"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		handleError(c, err)
	} else {
		if err := model.DeleteEvent(params.EventID); err != nil {
			handleError(c, err)
		} else {
			c.JSON(http.StatusOK, gin.H{"ok": true})
		}
	}
}

func handleError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{"error": err.Error()})
}
