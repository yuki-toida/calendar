package model

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/yuki-toida/knowme/config"
)

const emailDomain = "@candee.co.jp"

// User struct
type User struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Email     string    `json:"email" gorm:"not null"`
	Name      string    `json:"name"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Event struct
type Event struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Title     string    `json:"title" gorm:"not null"`
	StartDate time.Time `json:"startDate" gorm:"not null;type:date"`
	EndDate   time.Time `json:"endDate" gorm:"not null;type:date"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Migrate func
func Migrate() {
	db := config.ConnectDB()
	db.AutoMigrate(&User{}, &Event{})
}

// GetBase func
func GetBase(email, name, photo string) (User, []Event) {
	db := config.ConnectDB()
	user := GetUser(db, email)
	if email != "" && user == (User{}) {
		user.ID = strings.Replace(email, emailDomain, "", -1)
		user.Email = email
		user.Name = name
		user.Photo = photo
		db.Create(&user)
	}
	events := getEvents(db)
	return user, events
}

// GetUser func
func GetUser(db *gorm.DB, email string) User {
	var user User
	if email != "" {
		db.First(&user, "email = ?", email)
	}
	return user
}

func getEvents(db *gorm.DB) []Event {
	var events []Event
	db.Find(&events)
	return events
}

func getEvent(db *gorm.DB, id string) Event {
	var event Event
	db.First(&event, "id = ?", id)
	return event
}

// AddEvent func
func AddEvent(userID string, date time.Time) Event {
	db := config.ConnectDB()
	eventID := format(date) + ":" + userID
	event := getEvent(db, eventID)
	if event == (Event{}) {
		event.ID = eventID
		event.Title = userID
		event.StartDate = date
		event.EndDate = date
		db.Create(&event)
	}
	return event
}

// DeleteEvent func
func DeleteEvent(id string) {
	db := config.ConnectDB()
	event := getEvent(db, id)
	if event != (Event{}) {
		db.Delete(&event)
	}
}

func format(date time.Time) string {
	return date.Format("2006-01-02")
}
