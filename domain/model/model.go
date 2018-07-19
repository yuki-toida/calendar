package model

import (
	"time"
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
	Day       int       `gorm:"primary_key;type:int" json:"-"`
	Category  string    `gorm:"primary_key" json:"-"`
	ID        string    `gorm:"primary_key" json:"-"`
	EventID   string    `gorm:"unique;not null" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	StartDate time.Time `gorm:"type:date;not null" json:"startDate"`
	EndDate   time.Time `gorm:"type:date;not null" json:"endDate"`
	Ext       string    `gorm:"not null" json:"Ext"`
	Classes   string    `gorm:"-" json:"classes"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
