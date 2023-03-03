package controllers

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

var Version = "1.0.0"

func SetupHealthCheckController(router *mux.Router) {
	router.HandleFunc("/getVersion", getVersion).Methods("GET")
}

func getVersion(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, Version)
}
