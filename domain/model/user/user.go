package user

import (
	"time"

	"github.com/yuki-toida/knowme/domain/model"
)

// Event struct
type Event struct {
	Date     time.Time `json:"date"`
	Category string    `json:"category"`
	Titles   []string  `json:"titles"`
}

// Repository type
type Repository interface {
	First(id string) *model.User
	// Find(query interface{}, users []*model.User)
	Create(user *model.User)
	Update(user *model.User)
}

// Model type
type Model struct {
	Repository Repository
}

// New func
func New(r Repository) *Model {
	return &Model{
		Repository: r,
	}
}

func format(date time.Time) string {
	return date.Format("2006-01-02")
}

// SignIn func
func (m *Model) SignIn(id, name, photo string) *model.User {
	user := &model.User{
		ID:    id,
		Name:  name,
		Photo: photo,
	}
	if m.Repository.First(id) == nil {
		m.Repository.Create(user)
	} else {
		m.Repository.Update(user)
	}
	return user
}

// GetEvent func
func GetEvent(user *model.User) *Event {
	if user == nil {
		return nil
	}
	return nil
	// now := time.Now()
	// myEvent := getEvent(now.Year(), int(now.Month()), user.ID)
	// if myEvent == nil {
	// 	return nil
	// }
	// var events []Event
	// DB.Where(&Event{StartDate: myEvent.StartDate, Category: myEvent.Category}).Find(&events)
	// return &Event{
	// 	Date:     myEvent.StartDate,
	// 	Category: myEvent.Category,
	// 	Titles:   getEventTitles(events, myEvent.StartDate, myEvent.Category),
	// }
}

// GetEvents func
func GetEvents(id string) (*model.User, []Event) {
	return nil, []Event{}
	// allEvents := GetAllEvents()
	// myEvents := []Event{}
	// for _, v := range allEvents {
	// 	if v.ID == id {
	// 		myEvents = append(myEvents, v)
	// 	}
	// }
	// events := []Event{}
	// for _, v := range myEvents {
	// 	userEvent := Event{
	// 		Date:     v.StartDate,
	// 		Category: v.Category,
	// 		Titles:   getEventTitles(allEvents, v.StartDate, v.Category),
	// 	}
	// 	events = append(events, userEvent)
	// }
	// return nil, events
}
