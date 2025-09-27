package router

import (
	"go-gin-gorm-api/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(tagsController *controller.TagsController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello, World!")
	})

	tagRouter := router.Group("/tags")
	{
		tagRouter.POST("", tagsController.Create)
		tagRouter.PUT("/:tagId", tagsController.Update)
		tagRouter.DELETE("/:tagId", tagsController.Delete)
		tagRouter.GET("/:tagId", tagsController.FindById)
		tagRouter.GET("", tagsController.FindAll)
	}
	return router
}
