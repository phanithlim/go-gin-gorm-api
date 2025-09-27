package controller

import (
	req "go-gin-gorm-api/data/request"
	res "go-gin-gorm-api/data/response"
	"go-gin-gorm-api/helper"
	"go-gin-gorm-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagsController struct {
	tagService service.TagsService
}

func NewTagsController(service service.TagsService) *TagsController {
	return &TagsController{
		tagService: service,
	}
}

func (c *TagsController) Create(ctx *gin.Context) {
	createTagRequest := req.CreateTagsRequest{}
	err := ctx.ShouldBindJSON(&createTagRequest)
	helper.ErrorPanic(err)

	c.tagService.Create(createTagRequest)
	apiResponse := res.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, apiResponse)
}

func (c *TagsController) Update(ctx *gin.Context) {
	updateTagRequest := req.UpdateTagsRequest{}
	err := ctx.ShouldBindJSON(&updateTagRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	updateTagRequest.Id = id
	c.tagService.Update(updateTagRequest)
	apiResponse := res.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, apiResponse)
}

func (c *TagsController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	c.tagService.Delete(id)
	apiResponse := res.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, apiResponse)
}

func (c *TagsController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	tagResponse, err := c.tagService.FindById(id)
	helper.ErrorPanic(err)
	apiResponse := res.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tagResponse,
	}
	ctx.JSON(http.StatusOK, apiResponse)
}

func (c *TagsController) FindAll(ctx *gin.Context) {
	tagResponses := c.tagService.FindAll()
	apiResponse := res.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tagResponses,
	}
	ctx.JSON(http.StatusOK, apiResponse)
}
