package registry

import (
	"github.com/yuki-toida/knowme/domain/repository"
)

// Registry type
type Registry struct {
	UserRepository  repository.User
	EventRepository repository.Event
}

// NewRegistry func
func NewRegistry(u repository.User, e repository.Event) *Registry {
	return &Registry{
		UserRepository:  u,
		EventRepository: e,
	}
}
