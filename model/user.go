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

// Migrate func
func Migrate() {
	db := config.ConnectDB()
	db.AutoMigrate(&User{}, &Event{})
}

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
