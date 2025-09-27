package main

import (
	"go-gin-gorm-api/config"
	"go-gin-gorm-api/controller"
	"go-gin-gorm-api/helper"
	"go-gin-gorm-api/model"
	"go-gin-gorm-api/repository"
	"go-gin-gorm-api/router"
	"go-gin-gorm-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Starting the application...")
	db := config.DatabaseConnection()
	validate := validator.New()

	err := db.Table("tags").AutoMigrate(&model.Tags{})
	helper.ErrorPanic(err)
	log.Info().Msg("Database migration completed.")

	tagsRepository := repository.NewTagsRepository(db)
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)
	tagsController := controller.NewTagsController(tagsService)
	routes := router.NewRouter(tagsController)
	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err = server.ListenAndServe()
	helper.ErrorPanic(err)
}
