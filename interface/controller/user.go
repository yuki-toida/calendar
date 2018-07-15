package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuki-toida/knowme/domain/model/event"
	"github.com/yuki-toida/knowme/domain/model/user"
)

// UserEvents func
func UserEvents(c *gin.Context) {
	entity, _ := getUser(c)
	dayRest, nightRest := event.GetEventRest(time.Now())
	c.JSON(http.StatusOK, gin.H{
		"events":         event.GetEvents(entity),
		"myEvent":        user.GetEvent(entity),
		"dayEventRest":   dayRest,
		"nightEventRest": nightRest,
	})
}

// User func
func User(c *gin.Context) {
	id := c.Param("id")
	user, events := user.GetEvents(id)
	c.JSON(http.StatusOK, gin.H{
		"user":   user,
		"events": events,
	})
}
