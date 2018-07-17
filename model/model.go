package model

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/yuki-toida/knowme/domain/model"
)

func format(date time.Time) string {
	return date.Format("2006-01-02")
}

// GetUserEvent func
func GetUserEvent(user *model.User) *model.UserEvent {
	if user == nil {
		return nil
	}
	now := time.Now()
	myEvent := getEvent(now.Year(), int(now.Month()), user.ID)
	if myEvent == nil {
		return nil
	}
	var events []model.Event
	model.DB.Where(&model.Event{StartDate: myEvent.StartDate, Category: myEvent.Category}).Find(&events)
	return &model.UserEvent{
		Date:     myEvent.StartDate,
		Category: myEvent.Category,
		Titles:   getEventTitles(events, myEvent.StartDate, myEvent.Category),
	}
}

const eventCapacity = 3
const dayCouples = 4
const nightCouples = 8
const dayCategory = "day"
const nightCategory = "night"
const myDayClass = "text-white bg-danger rounded"
const dayClass = "text-white bg-danger rounded"
const myNightClass = "text-white bg-success rounded"
const nightClass = "text-white bg-success rounded"

func isDay(category string) bool {
	return category == dayCategory
}

// GetEvents func
func GetEvents(user *model.User) []model.Event {
	events := GetAllEvents()
	for i := range events {
		event := &events[i]
		if user != nil && event.ID == user.ID {
			if isDay(event.Category) {
				event.Classes = myDayClass
			} else {
				event.Classes = myNightClass
			}
		} else {
			if isDay(event.Category) {
				event.Classes = dayClass
			} else {
				event.Classes = nightClass
			}
		}
	}
	return events
}

// GetEventRest func
func GetEventRest(date time.Time) (int, int) {
	year := date.Year()
	month := int(date.Month())
	dayCount := count(year, month, dayCategory)
	nightCount := count(year, month, nightCategory)
	dayRest := dayCouples*eventCapacity - dayCount
	nightRest := nightCouples*eventCapacity - nightCount
	return dayRest, nightRest
}

// GetAllEvents func
func GetAllEvents() []model.Event {
	var events []model.Event
	model.DB.Find(&events)
	return events
}

// AddEvent func
func AddEvent(user *model.User, category string, date time.Time) (*model.Event, error) {
	year := date.Year()
	month := int(date.Month())
	if getEvent(year, month, user.ID) != nil {
		return nil, errors.New("今月は既に参加済みです")
	}
	if !verifyEventCapacity(year, month, date, category) {
		return nil, fmt.Errorf("定員（%d人）オーバーです", eventCapacity)
	}
	if !verifyCategoryCapacity(year, month, category) {
		return nil, errors.New("残席オーバーです")
	}
	if sameIDs := verifySameID(year, month, date, category, user.ID); 0 < len(sameIDs) {
		message := strings.Join(sameIDs, " ")
		return nil, errors.New(message + "と既に参加済みです")
	}
	var classes string
	if isDay(category) {
		classes = myDayClass
	} else {
		classes = myNightClass
	}
	event := &model.Event{
		Year:      date.Year(),
		Month:     int(date.Month()),
		ID:        user.ID,
		EventID:   format(date) + ":" + user.ID,
		Title:     user.Name,
		StartDate: date,
		EndDate:   date,
		Category:  category,
		Classes:   classes,
	}
	model.DB.Create(event)
	return event, nil
}

func getEvent(year, month int, id string) *model.Event {
	var event model.Event
	model.DB.Where(&model.Event{Year: year, Month: month, ID: id}).First(&event)
	if event == (model.Event{}) {
		return nil
	}
	return &event
}

func getEventTitles(events []model.Event, date time.Time, category string) []string {
	titles := []string{}
	for _, v := range events {
		if v.StartDate == date && v.Category == category {
			titles = append(titles, v.Title)
		}
	}
	return titles
}

func verifyEventCapacity(year, month int, date time.Time, category string) bool {
	var count int
	model.DB.Model(&model.Event{}).Where(&model.Event{Year: year, Month: month, StartDate: date, Category: category}).Count(&count)
	return count < eventCapacity
}

func verifyCategoryCapacity(year, month int, category string) bool {
	count := count(year, month, category)
	if isDay(category) {
		return count < dayCouples*eventCapacity
	}
	return count < nightCouples*eventCapacity
}

func verifySameID(year, month int, date time.Time, category, id string) []string {
	var yearEvents []model.Event
	model.DB.Where(&model.Event{Year: year}).Find(&yearEvents)
	myEvents := []model.Event{}
	for _, v := range yearEvents {
		if v.ID == id {
			myEvents = append(myEvents, v)
		}
	}
	ids := []string{}
	for _, x := range yearEvents {
		for _, y := range myEvents {
			if x.StartDate == y.StartDate && x.Category == y.Category {
				if x.ID != y.ID {
					ids = append(ids, x.ID)
				}
			}
		}
	}
	sameIDs := []string{}
	for _, x := range yearEvents {
		if x.StartDate == date && x.Category == category {
			for _, y := range ids {
				if x.ID == y {
					sameIDs = append(sameIDs, x.ID)
				}
			}
		}
	}
	return sameIDs
}

func count(year, month int, category string) int {
	var count int
	model.DB.Model(&model.Event{}).Where(&model.Event{Year: year, Month: month, Category: category}).Count(&count)
	return count
}
