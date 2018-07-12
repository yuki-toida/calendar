package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuki-toida/knowme/model"
)

// UserEvents func
func UserEvents(c *gin.Context) {
	user, _ := getUser(c)
	dayRest, nightRest := model.GetEventRest(time.Now())
	c.JSON(http.StatusOK, gin.H{
		"events":         model.GetEvents(user),
		"myEvent":        model.GetUserEvent(user),
		"dayEventRest":   dayRest,
		"nightEventRest": nightRest,
	})
}

// User func
func User(c *gin.Context) {
	id := c.Param("id")
	user, events := model.GetUserEvents(id)
	c.JSON(http.StatusOK, gin.H{
		"user":   user,
		"events": events,
	})
}
