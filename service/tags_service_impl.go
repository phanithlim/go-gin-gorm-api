package service

import (
	req "go-gin-gorm-api/data/request"
	res "go-gin-gorm-api/data/response"
	"go-gin-gorm-api/helper"
	"go-gin-gorm-api/model"
	"go-gin-gorm-api/repository"

	"github.com/go-playground/validator/v10"
)

type TagServiceImpl struct {
	TagsRepository repository.TagsRepository
	validate       *validator.Validate
}

func (t TagServiceImpl) Create(tags req.CreateTagsRequest) {
	err := t.validate.Struct(tags)
	helper.ErrorPanic(err)
	tagModel := model.Tags{
		Name: tags.Name,
	}
	t.TagsRepository.Save(tagModel)
}

func (t TagServiceImpl) Update(tags req.UpdateTagsRequest) {
	tagData, err := t.TagsRepository.FindById(tags.Id)
	helper.ErrorPanic(err)

	tagData.Name = tags.Name
	t.TagsRepository.Update(tagData)
}

func (t TagServiceImpl) Delete(id int) {
	err := t.validate.Struct(id)
	helper.ErrorPanic(err)
	t.TagsRepository.Delete(id)
}

func (t TagServiceImpl) FindById(id int) (tags res.TagsResponse, err error) {
	result, err := t.TagsRepository.FindById(id)
	helper.ErrorPanic(err)
	tags = res.TagsResponse{
		Id:   result.Id,
		Name: result.Name,
	}
	return tags, err
}

func (t TagServiceImpl) FindAll() []res.TagsResponse {
	result := t.TagsRepository.FindAll()
	tags := make([]res.TagsResponse, len(result))
	for i, v := range result {
		tags[i] = res.TagsResponse{
			Id:   v.Id,
			Name: v.Name,
		}
	}
	return tags
}

func NewTagsServiceImpl(tagsRepository repository.TagsRepository, validate *validator.Validate) TagsService {
	return &TagServiceImpl{
		TagsRepository: tagsRepository,
		validate:       validate,
	}
}
