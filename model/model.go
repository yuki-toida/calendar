package model

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/yuki-toida/knowme/config"
)

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

// Event struct
type Event struct {
	Year      int       `gorm:"primary_key;type:int" json:"-"`
	Month     int       `gorm:"primary_key;type:int" json:"-"`
	ID        string    `gorm:"primary_key" json:"-"`
	EventID   string    `gorm:"unique;not null" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	StartDate time.Time `gorm:"type:date;not null" json:"startDate"`
	EndDate   time.Time `gorm:"type:date;not null" json:"endDate"`
	Category  string    `gorm:"not null" json:"-"`
	Classes   string    `gorm:"-" json:"classes"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
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

func format(date time.Time) string {
	return date.Format("2006-01-02")
}

// SignIn func
func SignIn(id, name, photo string) (*User, error) {
	if id == "" || !strings.Contains(id, config.Config.Domain) {
		return nil, errors.New(config.Config.Domain + "を指定してください")
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
func GetEvents(user *User) []Event {
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
func GetAllEvents() []Event {
	var events []Event
	DB.Find(&events)
	return events
}

// AddEvent func
func AddEvent(user *User, category string, date time.Time) (*Event, error) {
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
	event := &Event{
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
	DB.Create(event)
	return event, nil
}

// DeleteEvent func
func DeleteEvent(user *User, category string, date time.Time) (*Event, error) {
	year := date.Year()
	month := int(date.Month())
	event := getEvent(year, month, user.ID)
	if event == nil {
		return nil, errors.New("参加していません")
	}
	DB.Delete(event)
	return event, nil
}

func getEvent(year, month int, id string) *Event {
	var event Event
	DB.Where(&Event{Year: year, Month: month, ID: id}).First(&event)
	if event == (Event{}) {
		return nil
	}
	return &event
}

func getEventTitles(events []Event, date time.Time, category string) []string {
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
	DB.Model(&Event{}).Where(&Event{Year: year, Month: month, StartDate: date, Category: category}).Count(&count)
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
	var yearEvents []Event
	DB.Where(&Event{Year: year}).Find(&yearEvents)
	myEvents := []Event{}
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
	DB.Model(&Event{}).Where(&Event{Year: year, Month: month, Category: category}).Count(&count)
	return count
}
