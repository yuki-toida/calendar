package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	service "github.com/yuki-toida/knowme/application/service/user"
	"github.com/yuki-toida/knowme/config"
	"github.com/yuki-toida/knowme/domain/model"
	"github.com/yuki-toida/knowme/interface/middleware"
)

func handleError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{"error": err.Error()})
}

func getUser(c *gin.Context) (*model.User, error) {
	return nil, errors.New("セッションユーザーが存在していません")
}

// HomeHealthz func
func HomeHealthz(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

// HomeIndex func
func HomeIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

// HomeInitial func
func HomeInitial(c *gin.Context) {
	id, _ := middleware.GetSessionID(c)
	registry := middleware.GetRegistry(c)
	user := service.Get(registry.UserRepository, id)
	fmt.Printf("%+v\n%+v\n%+v\n", id, registry, user)
	c.JSON(http.StatusOK, gin.H{
		"emailDomain": config.Config.EmailDomain,
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
		registry := middleware.GetRegistry(c)
		user, err := service.SignIn(registry.UserRepository, params.Email, params.Name, params.Photo)

		if err != nil {
			handleError(c, err)
		} else {
			middleware.SetSession(c, user.ID)
			c.JSON(http.StatusOK, gin.H{"user": user})
		}
	}
}

// HomeSignOut func
func HomeSignOut(c *gin.Context) {
	middleware.DeleteSession(c)
	c.JSON(http.StatusOK, gin.H{})
}
