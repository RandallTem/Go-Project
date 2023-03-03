package controllers

import (
	"CMSGo-backend/internal/cmsgo/dto"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"CMSGo-backend/internal/cmsgo/services"
)

var cmsService *services.CmsService

func SetupCmsController(router *mux.Router) {
	cmsService = services.GetCmsService()
	router.HandleFunc("/getApplications", getApplications).Methods("GET")
	router.HandleFunc("/findApplication", findApplication).Methods("POST")
	router.HandleFunc("/createApplication", createApplication).Methods("POST")
	router.HandleFunc("/deleteApplication", deleteApplication).Methods("POST")
}

func getApplications(response http.ResponseWriter, request *http.Request) {
	log.Println("getApplications() initiated")
	applications := cmsService.GetAllApplications()
	log.Printf("Found %d applications\n", len(*applications))
	responseJson, err := json.Marshal(applications)
	if err != nil {
		log.Println("Couldn't marshall applications to JSON")
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.Write(responseJson)
	log.Println("Sending response for getApplications()")
}

func findApplication(response http.ResponseWriter, request *http.Request) {
	var (
		error               bool
		applicationRequest  *dto.ApplicationRequest
		applicationResponse *dto.ApplicationResponse
	)
	log.Println("findApplication() initiated")
	applicationRequest, error = mapRequestToModel(request.Body)
	if error {
		applicationResponse = &dto.ApplicationResponse{
			IsSuccess:    false,
			ErrorMessage: "Error occurred during reading request body",
		}
		responseJson, _ := json.Marshal(applicationResponse)
		response.Header().Set("Content-Type", "application/json")
		response.Write(responseJson)
		return
	}
	log.Printf("Looking for application with uuid = %s, numauto = %s\n", applicationRequest.UUID, applicationRequest.Numauto)
	applicationResponse = cmsService.FindApplication(applicationRequest)
	if applicationResponse.IsSuccess {
		log.Println("Application was found")
	} else {
		log.Println("Application was not found")
	}
	responseJson, err := json.Marshal(applicationResponse)
	if err != nil {
		log.Println("Couldn't marshall application to JSON")
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.Write(responseJson)
	log.Println("Sending response to findApplication()")
}

func createApplication(response http.ResponseWriter, request *http.Request) {
	var (
		error               bool
		applicationRequest  *dto.ApplicationRequest
		applicationResponse *dto.ApplicationResponse
	)
	log.Println("createApplication() initiated")
	applicationRequest, error = mapRequestToModel(request.Body)
	if error {
		applicationResponse = &dto.ApplicationResponse{
			IsSuccess:    false,
			ErrorMessage: "Error occurred during reading request body",
		}
		responseJson, _ := json.Marshal(applicationResponse)
		response.Header().Set("Content-Type", "application/json")
		response.Write(responseJson)
		return
	}
	log.Printf("Creating new application with numauto = %s\n", applicationRequest.Numauto)
	applicationResponse = cmsService.CreateNewApplication(applicationRequest)
	if applicationResponse.IsSuccess {
		log.Printf("Application was created successfully with uuid = %s\n", applicationResponse.UUID)
	} else {
		log.Println("Couldn't create application")
	}
	responseJson, err := json.Marshal(applicationResponse)
	if err != nil {
		log.Println("Couldn't marshall application to JSON")
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.Write(responseJson)
	log.Println("Sending response to createApplication()")
}

func deleteApplication(response http.ResponseWriter, request *http.Request) {
	var (
		error               bool
		applicationRequest  *dto.ApplicationRequest
		applicationResponse *dto.ApplicationResponse
	)
	log.Println("deleteApplication() initiated")
	applicationRequest, error = mapRequestToModel(request.Body)
	if error {
		applicationResponse = &dto.ApplicationResponse{
			IsSuccess:    false,
			ErrorMessage: "Error occurred during reading request body",
		}
		responseJson, _ := json.Marshal(applicationResponse)
		response.Header().Set("Content-Type", "application/json")
		response.Write(responseJson)
		return
	}
	log.Printf("Deleting application with uuid = %s\n", applicationRequest.UUID)
	if applicationRequest.UUID == "" {
		log.Println("UUID is required to delete application")
		applicationResponse = &dto.ApplicationResponse{
			IsSuccess:    false,
			ErrorMessage: "UUID is empty",
		}
		responseJson, _ := json.Marshal(applicationResponse)
		response.Header().Set("Content-Type", "application/json")
		response.Write(responseJson)
		return
	}
	applicationResponse = cmsService.DeleteApplicationByUuid(applicationRequest.UUID)
	if applicationResponse.IsSuccess {
		log.Println("Application was deleted successfully")
	} else {
		log.Println("Couldn't delete application")
	}
	responseJson, _ := json.Marshal(applicationResponse)
	response.Header().Set("Content-Type", "application/json")
	response.Write(responseJson)
	log.Println("Sending response to deleteApplication()")
}

func mapRequestToModel(requestBody io.Reader) (*dto.ApplicationRequest, bool) {
	body, err := io.ReadAll(requestBody)
	if err != nil {
		log.Printf("Couldn't read request body\n")
		return nil, true
	}
	var applicationRequest dto.ApplicationRequest
	if err := json.Unmarshal(body, &applicationRequest); err != nil {
		log.Printf("Couldn't unmarshall request body")
		return nil, true
	}
	return &applicationRequest, false
}
