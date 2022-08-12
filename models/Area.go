package models

import "gorm.io/gorm"

type (
	Area struct {
		gorm.Model
		ID        int64  `gorm:"column:id;primaryKey;autoIncrement:true"`
		AreaValue int64  `gorm:"column:area_value"`
		AreaType  string `gorm:"column:type"`
	}
)
