package repository

import "github.com/yuki-toida/knowme/domain/model"

// User type
type User interface {
	First(string) *model.User
	Create(*model.User)
	Update(*model.User)
}

// Event type
type Event interface {
	First(int, int, string) *model.Event
	FindAll() []*model.Event
	Find(interface{}) []*model.Event
	Create(*model.Event)
	Delete(*model.Event)
}
