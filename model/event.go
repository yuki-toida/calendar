package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/yuki-toida/knowme/config"
)

const eventCapacity = 3
const myDayClass = "text-white bg-danger rounded"
const dayClass = "text-white bg-danger rounded"
const myNightClass = "text-white bg-primary rounded"
const nightClass = "text-white bg-primary rounded"

// Event struct
type Event struct {
	Year      int       `gorm:"primary_key;type:int" json:"-"`
	Month     int       `gorm:"primary_key;type:int" json:"-"`
	Category  string    `gorm:"primary_key" json:"-"`
	ID        string    `gorm:"primary_key" json:"-"`
	EventID   string    `gorm:"unique;not null" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	StartDate time.Time `gorm:"type:date;not null" json:"startDate"`
	EndDate   time.Time `gorm:"type:date;not null" json:"endDate"`
	Classes   string    `gorm:"-" json:"classes"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func isDay(category string) bool {
	return category == "day"
}

// GetEvents func
func GetEvents(user *User) []Event {
	db := config.ConnectDB()
	events := GetAllEvents(db)
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

// GetAllEvents func
func GetAllEvents(db *gorm.DB) []Event {
	var events []Event
	db.Find(&events)
	return events
}

// AddEvent func
func AddEvent(user *User, category string, date time.Time) (*Event, error) {
	db := config.ConnectDB()
	year := date.Year()
	month := int(date.Month())
	if getEvent(db, year, month, category, user.ID) != nil {
		return nil, errors.New("既に参加しています")
	}
	if !verifyEventCapacity(db, year, month, category, date) {
		return nil, fmt.Errorf("定員（%d人）オーバーです", eventCapacity)
	}
	var classes string
	if isDay(category) {
		classes = myDayClass
	} else {
		classes = myNightClass
	}
	event := &Event{
		Year:      date.Year(),
		Month:     int(date.Month()),
		Category:  category,
		ID:        user.ID,
		EventID:   format(date) + ":" + category + ":" + user.ID,
		Title:     user.Name,
		StartDate: date,
		EndDate:   date,
		Classes:   classes,
	}
	db.Create(event)
	return event, nil
}

func getEvent(db *gorm.DB, year, month int, category, id string) *Event {
	var event Event
	db.Where(&Event{Year: year, Month: month, Category: category, ID: id}).First(&event)
	if event == (Event{}) {
		return nil
	}
	return &event
}

func verifyEventCapacity(db *gorm.DB, year, month int, category string, date time.Time) bool {
	var count int
	db.Model(&Event{}).Where(&Event{Year: year, Month: month, Category: category, StartDate: date}).Count(&count)
	return count < eventCapacity
}

// DeleteEvent func
func DeleteEvent(user *User, category string, date time.Time) (*Event, error) {
	db := config.ConnectDB()
	year := date.Year()
	month := int(date.Month())
	event := getEvent(db, year, month, category, user.ID)
	if event == nil {
		return nil, errors.New("参加していません")
	}
	db.Delete(event)
	return event, nil
}
