package registry

import (
	"github.com/yuki-toida/knowme/domain/model/user"
)

// Registry type
type Registry struct {
	UserRepository user.Repository
}
