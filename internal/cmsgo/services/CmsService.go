package services

import (
	"CMSGo-backend/internal/cmsgo/dto"
	"CMSGo-backend/internal/cmsgo/models"
	"CMSGo-backend/internal/cmsgo/utils"
	"github.com/google/uuid"
	"time"
)

type CmsService struct {
	dbManager *utils.DBManager
}

var CmsServiceInstance CmsService

func GetCmsService() *CmsService {
	if CmsServiceInstance.dbManager == nil {
		CmsServiceInstance = CmsService{
			dbManager: utils.GetDBManager(),
		}
	}
	return &CmsServiceInstance
}

func (cmsService *CmsService) GetAllApplications() *[]dto.ApplicationResponse {
	applications := cmsService.dbManager.GetAllApplications()
	var appResponse []dto.ApplicationResponse
	for _, application := range *applications {
		appResponse = append(appResponse, dto.ApplicationResponse{
			UUID:                application.UUID,
			Numauto:             application.Numauto,
			ApplicationDateTime: application.ApplicationDateTime,
			IsSuccess:           true,
		})
	}
	return &appResponse
}

func (cmsService *CmsService) FindApplication(appReq *dto.ApplicationRequest) *dto.ApplicationResponse {
	var appResponse dto.ApplicationResponse
	var (
		fieldName  string
		fieldValue string
	)
	switch {
	case appReq.UUID != "":
		fieldName = "uuid"
		fieldValue = appReq.UUID
	case appReq.Numauto != "":
		fieldName = "numauto"
		fieldValue = appReq.Numauto
	default:
		return &dto.ApplicationResponse{}
	}
	application := cmsService.dbManager.GetApplicationByField(fieldName, fieldValue)
	appResponse = dto.ApplicationResponse{
		UUID:                application.UUID,
		Numauto:             application.Numauto,
		ApplicationDateTime: application.ApplicationDateTime,
		IsSuccess:           true,
	}
	if appResponse.UUID == "" {
		appResponse.IsSuccess = false
		appResponse.ApplicationDateTime = time.Now()
		appResponse.ErrorMessage = "Application was not found"
	}
	return &appResponse
}

func (cmsService *CmsService) CreateNewApplication(applicationRequest *dto.ApplicationRequest) *dto.ApplicationResponse {
	uuid := uuid.New().String()
	application := models.Application{
		UUID:                uuid,
		Numauto:             applicationRequest.Numauto,
		Application:         "{something is here}",
		ApplicationDateTime: time.Now(),
	}
	if len(application.Numauto) > 8 {
		application.Numauto = application.Numauto[:8]
	}
	cmsService.dbManager.SaveApplication(&application)
	application = *cmsService.dbManager.GetApplicationByField("uuid", uuid)
	if application.UUID == "" {
		return &dto.ApplicationResponse{
			ApplicationDateTime: time.Now(),
			IsSuccess:           false,
			ErrorMessage:        "Couldn't save application",
		}
	}
	return &dto.ApplicationResponse{
		UUID:                application.UUID,
		Numauto:             application.Numauto,
		ApplicationDateTime: application.ApplicationDateTime,
		IsSuccess:           true,
	}
}

func (cmsService *CmsService) DeleteApplicationByUuid(uuid string) *dto.ApplicationResponse {
	cmsService.dbManager.DeleteApplicationByUuid(uuid)
	application := *cmsService.dbManager.GetApplicationByField("uuid", uuid)
	if application.UUID == "" {
		return &dto.ApplicationResponse{
			ApplicationDateTime: time.Now(),
			IsSuccess:           true,
		}
	}
	return &dto.ApplicationResponse{
		ApplicationDateTime: time.Now(),
		IsSuccess:           false,
		ErrorMessage:        "Couldn't delete application. It still presents in DB",
	}
}
