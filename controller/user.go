package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuki-toida/knowme/model"
)

// User func
func User(c *gin.Context) {
	id := c.Param("id")
	user, events := model.GetUserEvents(id)
	c.JSON(http.StatusOK, gin.H{
		"user":   user,
		"events": events,
	})
}
