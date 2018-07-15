package user

import (
	"github.com/jinzhu/gorm"

	"github.com/yuki-toida/knowme/domain/model"
	"github.com/yuki-toida/knowme/domain/model/user"
)

type repository struct {
	db *gorm.DB
}

// New Repository
func New(db *gorm.DB) user.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) First(id string) *model.User {
	if id == "" {
		return nil
	}
	var result *model.User
	r.db.Where(&model.User{ID: id}).First(result)
	return result
}

func (r *repository) Create(user *model.User) {
	r.db.Create(user)
}

func (r *repository) Update(user *model.User) {
	r.db.Save(user)
}
