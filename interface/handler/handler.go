package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuki-toida/knowme/config"
	"github.com/yuki-toida/knowme/interface/middleware/session"
	"github.com/yuki-toida/knowme/interface/registry"
	"github.com/yuki-toida/knowme/usecase/event"
	"github.com/yuki-toida/knowme/usecase/user"
)

// Handler type
type Handler struct {
	registry *registry.Registry
}

// NewHandler func
func NewHandler(r *registry.Registry) *Handler {
	return &Handler{
		registry: r,
	}
}

func handleError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{"error": err.Error()})
}

// Initial func
func (h *Handler) Initial(c *gin.Context) {
	id := session.GetID(c)
	uc := user.NewUseCase(h.registry.UserRepository, h.registry.EventRepository)
	user := uc.Get(id)
	c.JSON(http.StatusOK, gin.H{
		"domain":       config.Config.Domain,
		"storageUrl":   config.Config.Server.StorageURL,
		"couplesDay":   event.CouplesDay,
		"couplesNight": event.CouplesNight,
		"user":         user,
	})
}

// SignIn func
func (h *Handler) SignIn(c *gin.Context) {
	var params struct {
		Email string `json:"email"`
		Name  string `json:"name"`
		Photo string `json:"photo"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		handleError(c, err)
	} else {
		uc := user.NewUseCase(h.registry.UserRepository, h.registry.EventRepository)
		user, err := uc.SignIn(params.Email, params.Name, params.Photo)
		if err != nil {
			handleError(c, err)
		} else {
			session.SaveID(c, user.ID)
			c.JSON(http.StatusOK, gin.H{"user": user})
		}
	}
}

// SignOut func
func (h *Handler) SignOut(c *gin.Context) {
	session.Delete(c)
	c.JSON(http.StatusOK, gin.H{})
}

// Search func
func (h *Handler) Search(c *gin.Context) {
	id := c.Param("id")
	uc := user.NewUseCase(h.registry.UserRepository, h.registry.EventRepository)
	user, events := uc.Search(id)
	c.JSON(http.StatusOK, gin.H{
		"user":   user,
		"events": events,
	})
}

// Upload func
func (h *Handler) Upload(c *gin.Context) {
	uc := user.NewUseCase(h.registry.UserRepository, h.registry.EventRepository)
	file, _ := c.FormFile("file")
	c.SaveUploadedFile(file, file.Filename)
	err := uc.Upload(file.Filename)
	if err != nil {
		handleError(c, err)
	} else {
		c.JSON(http.StatusOK, gin.H{})
	}
}

// GetEvent func
func (h *Handler) GetEvent(c *gin.Context) {
	id := session.GetID(c)
	user := user.NewUseCase(h.registry.UserRepository, h.registry.EventRepository).Get(id)
	uc := event.NewUseCase(h.registry.EventRepository)
	dayRestCount, nightRestCount := uc.GetRestCounts()

	c.JSON(http.StatusOK, gin.H{
		"events":         uc.Gets(),
		"event":          uc.GetUserEvent(user),
		"dayRestCount":   dayRestCount,
		"nightRestCount": nightRestCount,
	})
}

// CreateEvent func
func (h *Handler) CreateEvent(c *gin.Context) {
	var params struct {
		Category string    `json:"category"`
		Date     time.Time `json:"date"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		handleError(c, err)
	} else {
		id := session.GetID(c)
		user := user.NewUseCase(h.registry.UserRepository, h.registry.EventRepository).Get(id)
		uc := event.NewUseCase(h.registry.EventRepository)
		event, err := uc.CreateEvent(user, params.Category, params.Date.In(time.Local))
		if err != nil {
			handleError(c, err)
		} else {
			c.JSON(http.StatusOK, gin.H{"event": event})
		}
	}
}

// DeleteEvent func
func (h *Handler) DeleteEvent(c *gin.Context) {
	var params struct {
		Category string    `json:"category"`
		Date     time.Time `json:"date"`
	}
	if err := c.ShouldBindJSON(&params); err != nil {
		handleError(c, err)
	} else {
		id := session.GetID(c)
		uc := event.NewUseCase(h.registry.EventRepository)
		event, err := uc.Delete(id, params.Category, params.Date.In(time.Local))
		if err != nil {
			handleError(c, err)
		} else {
			c.JSON(http.StatusOK, gin.H{"event": event})
		}
	}
}
