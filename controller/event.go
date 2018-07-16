package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuki-toida/knowme/model"
)

// EventAdd func
func EventAdd(c *gin.Context) {
	var params struct {
		Category string    `json:"category"`
		Date     time.Time `json:"date"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		handleError(c, err)
	} else {
		user, err := getUser(c)
		if err != nil {
			handleError(c, err)
		}
		event, err := model.AddEvent(user, params.Category, params.Date.In(time.Local))
		if err != nil {
			handleError(c, err)
		} else {
			c.JSON(http.StatusOK, gin.H{"event": event})
		}
	}
}

// EventDelete func
func EventDelete(c *gin.Context) {
	var params struct {
		Category string    `json:"category"`
		Date     time.Time `json:"date"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		handleError(c, err)
	} else {
		user, err := getUser(c)
		if err != nil {
			handleError(c, err)
		}
		event, err := model.DeleteEvent(user, params.Category, params.Date.In(time.Local))
		if err != nil {
			handleError(c, err)
		} else {
			c.JSON(http.StatusOK, gin.H{"event": event})
		}
	}
}
