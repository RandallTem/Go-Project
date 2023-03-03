package utils

import (
	"fmt"
	"gorm.io/gorm"
	"time"

	"CMSGo-backend/internal/cmsgo/configs"
	"CMSGo-backend/internal/cmsgo/models"
)

type DBManager struct {
	dbConnection *gorm.DB
}

var DBManagerInstance DBManager

func GetDBManager() *DBManager {
	if DBManagerInstance.dbConnection == nil {
		DBManagerInstance = DBManager{
			dbConnection: configs.GetDataBaseConnection(),
		}
	}
	return &DBManagerInstance
}

func (dbManager *DBManager) GetAllApplications() *[]models.Application {
	var appsResponse []models.Application
	dbManager.dbConnection.Where("application_date_time > ?", time.Date(2023, time.February, 15, 0, 0, 0, 0, time.UTC)).Find(&appsResponse)
	return &appsResponse
}

func (dbManager *DBManager) GetApplicationByField(fieldName string, fieldValue string) *models.Application {
	var appResponse models.Application
	dbManager.dbConnection.Where(fmt.Sprintf("%s = ?", fieldName), fieldValue).First(&appResponse)
	return &appResponse
}

func (dbManager *DBManager) SaveApplication(application *models.Application) {
	dbManager.dbConnection.Save(application)
}

func (dbManager *DBManager) DeleteApplicationByUuid(uuid string) {
	dbManager.dbConnection.Where("uuid = ?", uuid).Delete(&models.Application{})
}
