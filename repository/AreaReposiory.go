package repository

import (
	"rest-api/models"

	"gorm.io/gorm"
)

type areaRepository struct {
	DB *gorm.DB
}

type AreaRepository interface {
	InsertArea(param1 int64, param2 int64, Type string, ar models.Area) (models.Area, error)
}

func NewAreaRepository(db *gorm.DB) AreaRepository {
	return areaRepository{
		DB: db,
	}
}

func (_r areaRepository) GetAll() (areas []models.Area, err error) {
	err = _r.DB.Find(&areas).Error
	return areas, err

}

func (_r areaRepository) InsertArea(param1 int64, param2 int64, Type string, ar models.Area) (models.Area, error) {
	var err error
	var area int64
	area = 0

	switch Type {
	case "persegi panjang":
		area = param1 * param2
		ar.AreaValue = area
		ar.AreaType = "persegi panjang"
		err = _r.DB.Create(&ar).Error
	case "persegi":
		area = param1 * param2
		ar.AreaValue = area
		ar.AreaType = "persegi"
		err = _r.DB.Create(&ar).Error
	case "segitiga":
		var x float64 = 0.5 * float64(param1) * float64(param2)
		area = int64(x)
		ar.AreaValue = area
		ar.AreaType = "segitiga"
		err = _r.DB.Create(&ar).Error
	default:
		ar.AreaValue = 0
		ar.AreaType = "undefined data"
		err = _r.DB.Create(&ar).Error
	}
	return ar, err
}
