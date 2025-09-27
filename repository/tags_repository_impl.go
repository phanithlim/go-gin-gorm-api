package repository

import (
	"errors"
	"go-gin-gorm-api/data/request"
	"go-gin-gorm-api/helper"
	"go-gin-gorm-api/model"

	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

func (t TagsRepositoryImpl) Save(tags model.Tags) {
	result := t.Db.Create(&tags)
	helper.ErrorPanic(result.Error)
}

func (t TagsRepositoryImpl) FindById(id int) (tags model.Tags, err error) {
	var tag model.Tags
	result := t.Db.Where("id = ?", id).Find(&tag)
	if result != nil {
		return tag, nil
	}
	return tag, errors.New("not found")
}

func (t TagsRepositoryImpl) Update(tags model.Tags) {
	var updateTags = request.UpdateTagsRequest{
		Id:   tags.Id,
		Name: tags.Name,
	}
	result := t.Db.Model(&tags).Where("id = ?", tags.Id).Updates(updateTags)
	helper.ErrorPanic(result.Error)
}

func (t TagsRepositoryImpl) Delete(id int) {
	var tags model.Tags
	result := t.Db.Where("id = ?", id).Delete(&tags)
	helper.ErrorPanic(result.Error)
}

func (t TagsRepositoryImpl) FindAll() []model.Tags {
	var tags []model.Tags
	result := t.Db.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags
}

func NewTagsRepository(db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{Db: db}
}
