package main

import (
	"rest-api/config"
	"rest-api/controllers"
	"rest-api/models"
	"rest-api/repository"
	"rest-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := config.OpenSqlite()
	db.AutoMigrate(&models.Area{})

	repo := repository.NewAreaRepository(db)
	service := services.NewAreaService(repo)
	controller := controllers.NewAreaController(service)

	api := r.Group("/api")
	api.POST("/area", controller.AddArea)

	r.Run()
}
