package services

import (
	"rest-api/models"
	"rest-api/repository"
	"rest-api/request"
)

type AreaService interface {
	InsertArea(request.AreaReq) (models.Area, error)
}

type areaService struct {
	areaRepository repository.AreaRepository
}

func NewAreaService(_s repository.AreaRepository) AreaService {
	return areaService{
		areaRepository: _s,
	}
}

func (_s areaService) InsertArea(area request.AreaReq) (models.Area, error) {
	return _s.areaRepository.InsertArea(area.Width, area.Height, area.Type, models.Area{})
}
