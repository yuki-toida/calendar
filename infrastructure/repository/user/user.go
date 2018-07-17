package user

import (
	"github.com/jinzhu/gorm"
	"github.com/yuki-toida/knowme/domain/model"
	"github.com/yuki-toida/knowme/domain/repository"
)

type repositoryImpl struct {
	db *gorm.DB
}

// New func
func New(db *gorm.DB) repository.User {
	return &repositoryImpl{
		db: db,
	}
}

func (r *repositoryImpl) First(id string) *model.User {
	result := model.User{}
	r.db.Where(&model.User{ID: id}).First(&result)
	if result == (model.User{}) {
		return nil
	}
	return &result
}

func (r *repositoryImpl) Create(u *model.User) {
	r.db.Create(u)
}

func (r *repositoryImpl) Update(u *model.User) {
	r.db.Save(u)
}
