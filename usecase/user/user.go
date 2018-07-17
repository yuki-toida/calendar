package user

import (
	"errors"
	"strings"

	"github.com/yuki-toida/knowme/config"
	"github.com/yuki-toida/knowme/domain/model"
	"github.com/yuki-toida/knowme/domain/repository"
)

// UseCase type
type UseCase struct {
	UserRepository  repository.User
	EventRepository repository.Event
}

// New func
func New(u repository.User, e repository.Event) *UseCase {
	return &UseCase{
		UserRepository:  u,
		EventRepository: e,
	}
}

// Get func
func (u *UseCase) Get(id string) *model.User {
	if id == "" {
		return nil
	}
	return u.UserRepository.First(id)
}

// SignIn func
func (u *UseCase) SignIn(email, name, photo string) (*model.User, error) {
	if email == "" || !strings.Contains(email, config.Config.Domain) {
		return nil, errors.New(config.Config.Domain + "を指定してください")
	}
	user := &model.User{
		ID:    email,
		Name:  name,
		Photo: photo,
	}
	if u.Get(email) == nil {
		model.DB.Create(user)
	} else {
		model.DB.Save(user)
	}
	return user, nil
}

// Search func
func (u *UseCase) Search(id string) (*model.User, []*model.UserEvent) {
	all := u.EventRepository.FindAll()
	myEvents := []*model.Event{}
	for _, v := range all {
		if v.ID == id {
			myEvents = append(myEvents, v)
		}
	}
	events := []*model.UserEvent{}
	for _, v := range myEvents {
		titles := []string{}
		for _, w := range all {
			if v.StartDate == w.StartDate && v.Category == w.Category {
				titles = append(titles, w.Title)
			}
		}
		events = append(events, &model.UserEvent{
			Date:     v.StartDate,
			Category: v.Category,
			Titles:   titles,
		})
	}
	return u.Get(id), events
}
