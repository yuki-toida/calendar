package session

import (
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yuki-toida/knowme/model"
)

const name = "id"

// AddMiddleware func
func AddMiddleware(router *gin.Engine) {
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("_knowme", store))
	router.Use(func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.URL.Path, "/static") {
			session := sessions.Default(c)
			if id := session.Get(name); id != nil {
				user := model.GetUser(id.(string))
				c.Set(name, user)
			}
		}
	})
}

// Get func
func Get(c *gin.Context) *model.User {
	user, exists := c.Get(name)
	if !exists {
		return nil
	}
	return user.(*model.User)
}

// Save func
func Save(c *gin.Context, id string) {
	session := sessions.Default(c)
	session.Set(name, id)
	session.Save()
}
