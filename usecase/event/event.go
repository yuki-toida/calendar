package event

import (
	"errors"
	"time"

	"github.com/yuki-toida/knowme/domain/model"
	"github.com/yuki-toida/knowme/domain/repository"
)

// UseCase type
type UseCase struct {
	EventRepository repository.Event
}

// New func
func New(u repository.Event) *UseCase {
	return &UseCase{
		EventRepository: u,
	}
}

// Get func
func (u *UseCase) Get(year, month int, id string) *model.Event {
	return u.EventRepository.First(year, month, id)
}

// Delete func
func (u *UseCase) Delete(id, category string, date time.Time) (*model.Event, error) {
	year := date.Year()
	month := int(date.Month())
	event := u.Get(year, month, id)
	if event == nil {
		return nil, errors.New("参加していません")
	}
	u.EventRepository.Delete(event)
	return event, nil
}
