package event

import (
	"github.com/jinzhu/gorm"
	"github.com/yuki-toida/knowme/domain/model"
	"github.com/yuki-toida/knowme/domain/repository"
)

type repositoryImpl struct {
	db *gorm.DB
}

// New func
func New(db *gorm.DB) repository.Event {
	return &repositoryImpl{
		db: db,
	}
}

func (r *repositoryImpl) First(year, month int, id string) *model.Event {
	result := model.Event{}
	r.db.Where(&model.Event{Year: year, Month: month, ID: id}).First(&result)
	if result == (model.Event{}) {
		return nil
	}
	return &result
}

func (r *repositoryImpl) FindAll() []*model.Event {
	results := []*model.Event{}
	r.db.Find(&results)
	return results
}

func (r *repositoryImpl) Find(query interface{}) []*model.Event {
	results := []*model.Event{}
	r.db.Where(query).Find(&results)
	return results
}

func (r *repositoryImpl) Create(u *model.Event) {
	r.db.Create(u)
}

func (r *repositoryImpl) Delete(u *model.Event) {
	r.db.Delete(u)
}
