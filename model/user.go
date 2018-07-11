package model

import (
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/yuki-toida/knowme/config"
)

// EmailDomain const
const EmailDomain = "@candee.co.jp"

func format(date time.Time) string {
	return date.Format("2006-01-02")
}

// User struct
type User struct {
	ID        string    `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// UserEvent struct
type UserEvent struct {
	Date     time.Time `json:"date"`
	Category string    `json:"category"`
	Titles   []string  `json:"titles"`
}

// SignIn func
func SignIn(id, name, photo string) (*User, error) {
	if id == "" || !strings.Contains(id, EmailDomain) {
		return nil, errors.New(EmailDomain + "を指定してください")
	}
	user := &User{
		ID:    id,
		Name:  name,
		Photo: photo,
	}
	db := config.ConnectDB()
	if GetUser(db, id) == nil {
		db.Create(user)
	} else {
		db.Save(user)
	}
	return user, nil
}

// GetUserEvent func
func GetUserEvent(user *User) *UserEvent {
	if user == nil {
		return nil
	}
	db := config.ConnectDB()
	now := time.Now()
	myEvent := getEvent(db, now.Year(), int(now.Month()), user.ID)
	if myEvent == nil {
		return nil
	}
	var events []Event
	db.Where(&Event{StartDate: myEvent.StartDate, Category: myEvent.Category}).Find(&events)
	return &UserEvent{
		Date:     myEvent.StartDate,
		Category: myEvent.Category,
		Titles:   getTitles(db, events, myEvent.StartDate, myEvent.Category),
	}
}

// GetUserEvents func
func GetUserEvents(id string) (*User, []UserEvent) {
	db := config.ConnectDB()
	allEvents := GetAllEvents(db)
	myEvents := []Event{}
	for _, v := range allEvents {
		if v.ID == id {
			myEvents = append(myEvents, v)
		}
	}
	events := []UserEvent{}
	for _, v := range myEvents {
		userEvent := UserEvent{
			Date:     v.StartDate,
			Category: v.Category,
			Titles:   getTitles(db, allEvents, v.StartDate, v.Category),
		}
		events = append(events, userEvent)
	}
	return GetUser(db, id), events
}

func getTitles(db *gorm.DB, events []Event, date time.Time, category string) []string {
	titles := []string{}
	for _, v := range events {
		if v.StartDate == date && v.Category == category {
			titles = append(titles, v.Title)
		}
	}
	return titles
}

// GetUser func
func GetUser(db *gorm.DB, id string) *User {
	if id == "" {
		return nil
	}
	var user User
	db.Where(&User{ID: id}).First(&user)
	if user == (User{}) {
		return nil
	}
	return &user
}
