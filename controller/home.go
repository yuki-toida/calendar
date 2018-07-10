package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/yuki-toida/knowme/config"
	"github.com/yuki-toida/knowme/model"
)

func handleError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{"error": err.Error()})
}

func getUser(c *gin.Context) (*model.User, error) {
	user, exists := c.Get(config.SessionName)
	if exists {
		return user.(*model.User), nil
	}
	return nil, errors.New("セッションユーザーが存在していません")
}

// HomeHealthz func
func HomeHealthz(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

// HomeIndex func
func HomeIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"staticPath": config.Config.Server.StaticPath})
}

// HomeInitial func
func HomeInitial(c *gin.Context) {
	user, _ := getUser(c)
	c.JSON(http.StatusOK, gin.H{
		"staticPath":  config.Config.Server.StaticPath,
		"emailDomain": model.EmailDomain,
		"events":      model.GetEvents(user),
		"user":        user,
	})
}

// HomeSignIn func
func HomeSignIn(c *gin.Context) {
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
			session.Set(config.SessionName, user.ID)
			session.Save()
			c.JSON(http.StatusOK, gin.H{"user": user})
		}
	}
}

// HomeSignOut func
func HomeSignOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{})
}
