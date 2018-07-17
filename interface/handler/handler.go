package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuki-toida/knowme/config"
	"github.com/yuki-toida/knowme/interface/middleware/session"
	"github.com/yuki-toida/knowme/interface/registry"
	"github.com/yuki-toida/knowme/model"
	"github.com/yuki-toida/knowme/usecase/event"
	"github.com/yuki-toida/knowme/usecase/user"
)

func handleError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{"error": err.Error()})
}

// GETInitial func
func GETInitial(c *gin.Context, r *registry.Registry) {
	id := session.GetID(c)
	uc := user.New(r.UserRepository, r.EventRepository)
	user := uc.Get(id)
	c.JSON(http.StatusOK, gin.H{
		"domain": config.Config.Domain,
		"user":   user,
	})
}

// POSTSignIn func
func POSTSignIn(c *gin.Context, r *registry.Registry) {
	var params struct {
		Email string `json:"email"`
		Name  string `json:"name"`
		Photo string `json:"photo"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		handleError(c, err)
	} else {
		uc := user.New(r.UserRepository, r.EventRepository)
		user, err := uc.SignIn(params.Email, params.Name, params.Photo)
		if err != nil {
			handleError(c, err)
		} else {
			session.SaveID(c, user.ID)
			c.JSON(http.StatusOK, gin.H{"user": user})
		}
	}
}

// DELETESignOut func
func DELETESignOut(c *gin.Context, r *registry.Registry) {
	session.Delete(c)
	c.JSON(http.StatusOK, gin.H{})
}

// GETSearch func
func GETSearch(c *gin.Context, r *registry.Registry) {
	id := c.Param("id")
	uc := user.New(r.UserRepository, r.EventRepository)
	user, events := uc.Search(id)
	c.JSON(http.StatusOK, gin.H{
		"user":   user,
		"events": events,
	})
}

// GETEvent func
func GETEvent(c *gin.Context, r *registry.Registry) {
	id := session.GetID(c)
	uc := user.New(r.UserRepository, r.EventRepository)
	user := uc.Get(id)

	dayRest, nightRest := model.GetEventRest(time.Now())
	c.JSON(http.StatusOK, gin.H{
		"events":         model.GetEvents(user),
		"myEvent":        model.GetUserEvent(user),
		"dayEventRest":   dayRest,
		"nightEventRest": nightRest,
	})
}

// POSTEvent func
func POSTEvent(c *gin.Context, r *registry.Registry) {
	var params struct {
		Category string    `json:"category"`
		Date     time.Time `json:"date"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		handleError(c, err)
	} else {
		id := session.GetID(c)
		uc := user.New(r.UserRepository, r.EventRepository)
		user := uc.Get(id)
		event, err := model.AddEvent(user, params.Category, params.Date.In(time.Local))
		if err != nil {
			handleError(c, err)
		} else {
			c.JSON(http.StatusOK, gin.H{"event": event})
		}
	}
}

// PUTEvent func
func PUTEvent(c *gin.Context, r *registry.Registry) {
	var params struct {
		Category string    `json:"category"`
		Date     time.Time `json:"date"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		handleError(c, err)
	} else {
		id := session.GetID(c)
		uc := event.New(r.EventRepository)
		event, err := uc.Delete(id, params.Category, params.Date.In(time.Local))
		if err != nil {
			handleError(c, err)
		} else {
			c.JSON(http.StatusOK, gin.H{"event": event})
		}
	}
}
