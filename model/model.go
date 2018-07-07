package model

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/yuki-toida/knowme/config"
)

// EmailDomain const
const EmailDomain = "@candee.co.jp"

// User struct
type User struct {
	UserID    string    `gorm:"primary_key" json:"id"`
	Email     string    `gorm:"not null" json:"email"`
	Name      string    `json:"name"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Event struct
type Event struct {
	Year      int       `gorm:"primary_key;type:int" json:"-"`
	Month     int       `gorm:"primary_key;type:int" json:"-"`
	UserID    string    `gorm:"primary_key" json:"-"`
	EventID   string    `gorm:"unique;not null" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	StartDate time.Time `gorm:"type:date;not null" json:"startDate"`
	EndDate   time.Time `gorm:"type:date;not null" json:"endDate"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Migrate func
func Migrate() {
	db := config.ConnectDB()
	db.AutoMigrate(&User{}, &Event{})
}

// InitUser func
func InitUser(email, name, photo string) User {
	db := config.ConnectDB()
	userID := strings.Replace(email, EmailDomain, "", -1)
	user := GetUser(db, userID)
	if user == (User{}) {
		user.UserID = userID
		user.Email = email
		user.Name = name
		user.Photo = photo
		db.Create(&user)
	}
	return user
}

// GetUser func
func GetUser(db *gorm.DB, userID string) User {
	var user User
	if userID != "" {
		db.Where(&User{UserID: userID}).First(&user)
	}
	return user
}

// GetEvents func
func GetEvents() []Event {
	db := config.ConnectDB()
	var events []Event
	db.Find(&events)
	return events
}

// AddEvent func
func AddEvent(userID string, date time.Time) (Event, error) {
	db := config.ConnectDB()
	user := GetUser(db, userID)
	if user == (User{}) {
		return Event{}, errors.New("Invalid userID")
	}
	var event Event
	db.Where(&Event{Year: date.Year(), Month: int(date.Month()), UserID: userID}).First(&event)
	fmt.Println(event)
	if event != (Event{}) {
		return Event{}, errors.New("Event exists")
	}
	event.Year = date.Year()
	event.Month = int(date.Month())
	event.UserID = userID
	event.EventID = format(date) + ":" + userID
	event.Title = user.Name
	event.StartDate = date
	event.EndDate = date
	db.Create(&event)
	return event, nil
}

// DeleteEvent func
func DeleteEvent(eventID string) error {
	db := config.ConnectDB()
	var event Event
	db.Where(&Event{EventID: eventID}).First(&event)
	if event == (Event{}) {
		return errors.New("Event doesn't exist")
	}
	db.Delete(&event)
	return nil
}

func format(date time.Time) string {
	return date.Format("2006-01-02")
}
