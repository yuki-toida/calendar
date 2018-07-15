package model

import "time"

// User struct
type User struct {
	ID        string    `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

// Event struct
type Event struct {
	Year      int       `gorm:"primary_key;type:int" json:"-"`
	Month     int       `gorm:"primary_key;type:int" json:"-"`
	ID        string    `gorm:"primary_key" json:"-"`
	EventID   string    `gorm:"not null" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	StartDate time.Time `gorm:"type:date;not null" json:"startDate"`
	EndDate   time.Time `gorm:"type:date;not null" json:"endDate"`
	Category  string    `gorm:"not null" json:"-"`
	Classes   string    `gorm:"-" json:"classes"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
