package session

import (
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
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
				c.Set(name, id)
			}
		}
	})
}

// GetID func
func GetID(c *gin.Context) string {
	user, exists := c.Get(name)
	if !exists {
		return ""
	}
	return user.(string)
}

// SaveID func
func SaveID(c *gin.Context, id string) {
	session := sessions.Default(c)
	session.Set(name, id)
	session.Save()
}

// Delete func
func Delete(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}
