package service

import (
	req "go-gin-gorm-api/data/request"
	res "go-gin-gorm-api/data/response"
)

type TagsService interface {
	Create(tags req.CreateTagsRequest)
	Update(tags req.UpdateTagsRequest)
	Delete(id int)
	FindById(id int) (tags res.TagsResponse, err error)
	FindAll() []res.TagsResponse
}
