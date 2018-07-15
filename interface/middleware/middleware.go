package middleware

import (
	"errors"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/yuki-toida/knowme/registry"
)

const registryName = "registry"
const sessionName = "id"

// AddRegistry func
func AddRegistry(r registry.Registry) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.URL.Path, "/static") {
			c.Set(registryName, r)
		}
	}
}

// GetRegistry func
func GetRegistry(c *gin.Context) registry.Registry {
	return c.MustGet(registryName).(registry.Registry)
}

// AddSession func
func AddSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.URL.Path, "/static") {
			session := sessions.Default(c)
			if id := session.Get(sessionName); id != nil {
				c.Set(sessionName, id.(string))
			}
		}
	}
}

// SetSession func
func SetSession(c *gin.Context, id string) {
	session := sessions.Default(c)
	session.Set(sessionName, id)
	session.Save()
}

// GetSessionID func
func GetSessionID(c *gin.Context) (string, error) {
	id, exists := c.Get(sessionName)
	if !exists {
		return "", errors.New("session doesn't exist")
	}
	return id.(string), nil
}

// DeleteSession func
func DeleteSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}
