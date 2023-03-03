package models

import "time"

type Application struct {
	UUID                string    `gorm:"column:uuid;primaryKey"`
	Numauto             string    `gorm:"colunn:numauto"`
	Application         string    `gorm:"column:application"`
	ApplicationDateTime time.Time `gorm:"column:application_date_time"`
}

func (Application) TableName() string {
	return "cms.applications"
}
