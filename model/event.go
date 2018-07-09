package model

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/yuki-toida/knowme/config"
)

// EmailDomain const
const myEventClass = "text-white bg-primary rounded"
const otherEventClass = "rounded"

// Event struct
type Event struct {
	Year      int       `gorm:"primary_key;type:int" json:"-"`
	Month     int       `gorm:"primary_key;type:int" json:"-"`
	ID        string    `gorm:"primary_key" json:"-"`
	EventID   string    `gorm:"unique;not null" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	StartDate time.Time `gorm:"type:date;not null" json:"startDate"`
	EndDate   time.Time `gorm:"type:date;not null" json:"endDate"`
	Classes   string    `gorm:"-" json:"classes"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// GetEvents func
func GetEvents(user *User) []Event {
	db := config.ConnectDB()
	var events []Event
	db.Find(&events)
	for i := range events {
		event := &events[i]
		if user != nil && event.ID == user.ID {
			event.Classes = myEventClass
		} else {
			event.Classes = otherEventClass
		}
	}
	return events
}

// AddEvent func
func AddEvent(user *User, date time.Time) (*Event, error) {
	db := config.ConnectDB()
	if anyEvent(db, date.Year(), int(date.Month()), user.ID) {
		return nil, errors.New("既にイベントに参加しています")
	}
	event := &Event{
		Year:      date.Year(),
		Month:     int(date.Month()),
		ID:        user.ID,
		EventID:   format(date) + ":" + user.ID,
		Title:     user.Name,
		StartDate: date,
		EndDate:   date,
		Classes:   myEventClass,
	}
	db.Create(event)
	return event, nil
}

func anyEvent(db *gorm.DB, year, month int, id string) bool {
	var event Event
	db.Where(&Event{Year: year, Month: month, ID: id}).First(&event)
	return event != (Event{})
}

// DeleteEvent func
func DeleteEvent(user *User, eventID string) error {
	db := config.ConnectDB()
	var event Event
	db.Where(&Event{EventID: eventID}).First(&event)
	if event == (Event{}) {
		return errors.New("削除するイベントが存在していません")
	}
	if event.ID != user.ID {
		return errors.New("自分のイベントではありません")
	}
	db.Delete(&event)
	return nil
}
