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

// DB var
var DB *gorm.DB

// Initialize func
func Initialize() {
	connectionString := "root:zaqroot@tcp(" + config.Config.Db.Host + ":" + config.Config.Db.Port + ")/" + config.Config.Db.Name + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	db.LogMode(true)
	db.AutoMigrate(&User{}, &Event{})
	DB = db
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
	if GetUser(id) == nil {
		DB.Create(user)
	} else {
		DB.Save(user)
	}
	return user, nil
}

// GetUserEvent func
func GetUserEvent(user *User) *UserEvent {
	if user == nil {
		return nil
	}
	now := time.Now()
	myEvent := getEvent(now.Year(), int(now.Month()), user.ID)
	if myEvent == nil {
		return nil
	}
	var events []Event
	DB.Where(&Event{StartDate: myEvent.StartDate, Category: myEvent.Category}).Find(&events)
	return &UserEvent{
		Date:     myEvent.StartDate,
		Category: myEvent.Category,
		Titles:   getEventTitles(events, myEvent.StartDate, myEvent.Category),
	}
}

// GetUserEvents func
func GetUserEvents(id string) (*User, []UserEvent) {
	allEvents := GetAllEvents()
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
			Titles:   getEventTitles(allEvents, v.StartDate, v.Category),
		}
		events = append(events, userEvent)
	}
	return GetUser(id), events
}

// GetUser func
func GetUser(id string) *User {
	if id == "" {
		return nil
	}
	var user User
	DB.Where(&User{ID: id}).First(&user)
	if user == (User{}) {
		return nil
	}
	return &user
}
