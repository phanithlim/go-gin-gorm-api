package repository

import "go-gin-gorm-api/model"

type TagsRepository interface {
	Save(tags model.Tags)
	FindById(id int) (tags model.Tags, err error)
	Update(tags model.Tags)
	Delete(id int)
	FindAll() []model.Tags
}
