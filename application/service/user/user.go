package user

import (
	"errors"
	"strings"

	"github.com/yuki-toida/knowme/config"
	"github.com/yuki-toida/knowme/domain/model"
	"github.com/yuki-toida/knowme/domain/model/user"
)

// Get func
func Get(repository user.Repository, id string) *model.User {
	return repository.First(id)
}

// SignIn func
func SignIn(repository user.Repository, id, name, photo string) (*model.User, error) {
	if id == "" || !strings.Contains(id, config.Config.EmailDomain) {
		return nil, errors.New(config.Config.EmailDomain + "を指定してください")
	}
	model := user.New(repository)
	return model.SignIn(id, name, photo), nil
}
