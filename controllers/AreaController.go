package controllers

import (
	"net/http"
	"rest-api/request"
	"rest-api/services"

	"github.com/gin-gonic/gin"
)

type AreaController interface {
	AddArea(*gin.Context)
}

type areaController struct {
	service services.AreaService
}

func NewAreaController(_s services.AreaService) AreaController {
	return areaController{
		service: _s,
	}
}

func (_c areaController) AddArea(c *gin.Context) {
	var areaReq request.AreaReq

	if err := c.ShouldBindJSON(&areaReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error field can't be empty",
		})
		return
	}

	area, err := _c.service.InsertArea(areaReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error insert: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success insert area",
		"data":    area,
	})
}
