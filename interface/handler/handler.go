package handler

import (
	"net/http"
	"strconv"
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

// Pictures func
func (h *Handler) Pictures(c *gin.Context) {
	id := session.GetID(c)
	user := user.NewUseCase(h.registry.UserRepository, h.registry.EventRepository).Get(id)
	uc := event.NewUseCase(h.registry.EventRepository)
	c.JSON(http.StatusOK, gin.H{
		"pictures": uc.GetAllPictures(),
		"events":   uc.GetUserEvents(user),
	})
}

// Upload func
func (h *Handler) Upload(c *gin.Context) {
	id := session.GetID(c)
	uc := event.NewUseCase(h.registry.EventRepository)
	year, _ := strconv.Atoi(c.PostForm("year"))
	month, _ := strconv.Atoi(c.PostForm("month"))
	day, _ := strconv.Atoi(c.PostForm("day"))
	category := c.PostForm("category")
	file, _ := c.FormFile("file")

	c.SaveUploadedFile(file, file.Filename)
	err := uc.Upload(year, month, day, category, id, file.Filename)
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
	dayRestCouples, nightRestCouples := uc.GetRestCouples()

	c.JSON(http.StatusOK, gin.H{
		"events":           uc.Gets(),
		"event":            uc.GetUserEvent(user),
		"dayRestCouples":   dayRestCouples,
		"nightRestCouples": nightRestCouples,
		"pictures":         uc.GetPictures(),
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
		event, err := uc.CreateEvent(user, params.Category, params.Date.In(config.JST))
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
		event, err := uc.Delete(id, params.Category, params.Date.In(config.JST))
		if err != nil {
			handleError(c, err)
		} else {
			c.JSON(http.StatusOK, gin.H{"event": event})
		}
	}
}
