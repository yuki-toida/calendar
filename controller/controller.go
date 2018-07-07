package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/yuki-toida/knowme/config"
	"github.com/yuki-toida/knowme/model"
)

const sessionName = "UserID"

// Healthz func
func Healthz(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

// Index func
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"staticUrl": config.Config.Server.StaticURL})
}

// Initial func
func Initial(c *gin.Context) {
	session := sessions.Default(c)
	var user model.User
	if userID := session.Get(sessionName); userID != nil {
		user = model.GetUser(userID.(string))
	}
	fmt.Println(user)
	c.JSON(http.StatusOK, gin.H{
		"staticUrl":   config.Config.Server.StaticURL,
		"emailDomain": model.EmailDomain,
		"events":      model.GetEvents(),
		"user":        user,
	})
}

// SignIn func
func SignIn(c *gin.Context) {
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
			session := sessions.Default(c)
			session.Set(sessionName, user.UserID)
			session.Save()
			c.JSON(http.StatusOK, gin.H{"user": user})
		}
	}
}

// SignOut func
func SignOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{"user": model.User{}})
}

// AddEvent func
func AddEvent(c *gin.Context) {
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

// DeleteEvent func
func DeleteEvent(c *gin.Context) {
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
