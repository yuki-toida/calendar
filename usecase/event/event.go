package event

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/yuki-toida/knowme/domain/model"
	"github.com/yuki-toida/knowme/domain/repository"
)

// CouplesDay const
const CouplesDay = 4

// CouplesNight const
const CouplesNight = 8

const capacity = 3
const categoryDay = "day"
const categoryNight = "night"
const classDay = "text-white bg-danger rounded"
const classNight = "text-white bg-success rounded"

// UseCase type
type UseCase struct {
	EventRepository repository.Event
}

// NewUseCase func
func NewUseCase(u repository.Event) *UseCase {
	return &UseCase{
		EventRepository: u,
	}
}

// Get func
func (u *UseCase) Get(year, month, day int, category string, id string) *model.Event {
	return u.EventRepository.First(year, month, day, category, id)
}

// Gets func
func (u *UseCase) Gets() []*model.Event {
	events := u.EventRepository.FindAll()
	for _, v := range events {
		if v.Category == categoryDay {
			v.Classes = classDay
		} else {
			v.Classes = classNight
		}
	}
	return events
}

// GetUserEvent func
func (u *UseCase) GetUserEvent(user *model.User) *model.UserEvent {
	if user == nil {
		return nil
	}
	now := time.Now()
	year := now.Year()
	month := int(now.Month())

	events := u.EventRepository.Find(&model.Event{Year: year, Month: month, ID: user.ID})
	if len(events) <= 0 {
		return nil
	}
	event := events[0]
	titles := []string{}
	for _, v := range u.EventRepository.Find(&model.Event{StartDate: event.StartDate, Category: event.Category}) {
		if v.StartDate == event.StartDate && v.Category == event.Category {
			titles = append(titles, v.Title)
		}
	}

	return &model.UserEvent{
		Date:     event.StartDate,
		Category: event.Category,
		Titles:   titles,
	}
}

// GetRestCounts func
func (u *UseCase) GetRestCounts() (int, int) {
	now := time.Now()
	year := now.Year()
	month := int(now.Month())
	days := len(u.EventRepository.Find(&model.Event{Year: year, Month: month, Category: categoryDay}))
	nights := len(u.EventRepository.Find(&model.Event{Year: year, Month: month, Category: categoryNight}))
	dayRestCount := CouplesDay*capacity - days
	nightRestCount := CouplesNight*capacity - nights
	return dayRestCount, nightRestCount
}

// CreateEvent func
func (u *UseCase) CreateEvent(user *model.User, category string, date time.Time) (*model.Event, error) {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	if date.Before(today) {
		return nil, errors.New("過去の登録は出来ません")
	}
	year := date.Year()
	month := int(date.Month())
	day := date.Day()
	if today.Year() < year || int(today.Month()) < month {
		return nil, errors.New("未来の登録は出来ません")
	}
	if u.Get(year, month, day, category, user.ID) != nil {
		return nil, errors.New("今月は既に登録済みです")
	}
	dateCount := len(u.EventRepository.Find(&model.Event{Year: year, Month: month, StartDate: date, Category: category}))
	if capacity <= dateCount {
		return nil, errors.New("すでに満席です")
	}
	categoryCount := len(u.EventRepository.Find(&model.Event{Year: year, Month: month, Category: category}))
	switch category {
	case categoryDay:
		if CouplesDay*capacity <= categoryCount {
			return nil, errors.New("昼Knowmeはすでに満席です")
		}
	case categoryNight:
		if CouplesNight*capacity <= categoryCount {
			return nil, errors.New(category + "夜Knowmeはすでに満席です")
		}
	}

	if duplicateIds := u.duplicateIds(year, month, date, category, user.ID); 0 < len(duplicateIds) {
		message := strings.Join(duplicateIds, " ")
		return nil, errors.New(message + "と既に参加済みです")
	}
	var classes string
	if category == categoryDay {
		classes = classDay
	} else {
		classes = classNight
	}
	event := &model.Event{
		Year:      date.Year(),
		Month:     int(date.Month()),
		Day:       date.Day(),
		Category:  category,
		ID:        user.ID,
		EventID:   date.Format("2006-01-02") + ":" + user.ID,
		Title:     user.Name,
		StartDate: date,
		EndDate:   date,
		Classes:   classes,
	}
	u.EventRepository.Create(event)
	return event, nil
}

func (u *UseCase) duplicateIds(year, month int, date time.Time, category, id string) []string {
	years := u.EventRepository.Find(&model.Event{Year: year})
	events := []*model.Event{}
	for _, v := range years {
		if v.ID == id {
			events = append(events, v)
		}
	}
	duplicateIds := []string{}
	for _, x := range years {
		for _, y := range events {
			if x.StartDate == y.StartDate && x.Category == y.Category {
				if x.ID != y.ID {
					duplicateIds = append(duplicateIds, x.ID)
				}
			}
		}
	}
	results := []string{}
	for _, x := range years {
		if x.StartDate == date && x.Category == category {
			for _, y := range duplicateIds {
				if x.ID == y {
					results = append(results, x.ID)
				}
			}
		}
	}
	fmt.Println(duplicateIds)
	fmt.Println(results)
	return results
}

// Delete func
func (u *UseCase) Delete(id, category string, date time.Time) (*model.Event, error) {
	event := u.Get(date.Year(), int(date.Month()), date.Day(), category, id)
	if event == nil {
		return nil, errors.New("参加していません")
	}
	u.EventRepository.Delete(event)
	return event, nil
}
