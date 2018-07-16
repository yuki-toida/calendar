package handler

import (
	"net/http"
	"time"

	"github.com/yuki-toida/knowme/config"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yuki-toida/knowme/interface/middleware/session"
	"github.com/yuki-toida/knowme/model"
)

func handleError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{"error": err.Error()})
}

// GETInitial func
func GETInitial(c *gin.Context) {
	user := session.Get(c)
	c.JSON(http.StatusOK, gin.H{
		"domain": config.Config.Domain,
		"user":   user,
	})
}

// POSTSignIn func
func POSTSignIn(c *gin.Context) {
	var params struct {
		Email string `json:"email"`
		Name  string `json:"name"`
		Photo string `json:"photo"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		handleError(c, err)
	} else {
		user, err := model.SignIn(params.Email, params.Name, params.Photo)
		if err != nil {
			handleError(c, err)
		} else {
			session.Save(c, user.ID)
			c.JSON(http.StatusOK, gin.H{"user": user})
		}
	}
}

// DELETESignOut func
func DELETESignOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{})
}

// GETUserEvent func
func GETUserEvent(c *gin.Context) {
	user := session.Get(c)
	dayRest, nightRest := model.GetEventRest(time.Now())
	c.JSON(http.StatusOK, gin.H{
		"events":         model.GetEvents(user),
		"myEvent":        model.GetUserEvent(user),
		"dayEventRest":   dayRest,
		"nightEventRest": nightRest,
	})
}

// GETUserSearch func
func GETUserSearch(c *gin.Context) {
	id := c.Param("id")
	user, events := model.GetUserEvents(id)
	c.JSON(http.StatusOK, gin.H{
		"user":   user,
		"events": events,
	})
}

// POSTEvent func
func POSTEvent(c *gin.Context) {
	var params struct {
		Category string    `json:"category"`
		Date     time.Time `json:"date"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		handleError(c, err)
	} else {
		user := session.Get(c)
		event, err := model.AddEvent(user, params.Category, params.Date.In(time.Local))
		if err != nil {
			handleError(c, err)
		} else {
			c.JSON(http.StatusOK, gin.H{"event": event})
		}
	}
}

// PUTEvent func
func PUTEvent(c *gin.Context) {
	var params struct {
		Category string    `json:"category"`
		Date     time.Time `json:"date"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		handleError(c, err)
	} else {
		user := session.Get(c)
		event, err := model.DeleteEvent(user, params.Category, params.Date.In(time.Local))
		if err != nil {
			handleError(c, err)
		} else {
			c.JSON(http.StatusOK, gin.H{"event": event})
		}
	}
}
