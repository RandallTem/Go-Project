package cmsgo

import (
	"CMSGo-backend/internal/cmsgo/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"CMSGo-backend/internal/cmsgo/configs"
)

type CMSGoServer struct {
	serverConfig *configs.ServerConfig
	router       *http.Handler
}

func StartServer() error {
	server := configureServer()
	log.Printf("Starting server on port %s\n", server.serverConfig.ServerPort)
	return http.ListenAndServe(server.serverConfig.ServerPort, *server.router)
}

func configureServer() *CMSGoServer {
	serverConfig := configs.GetServerConfig()
	router := mux.NewRouter()
	controllers.SetupCmsController(router)
	controllers.SetupHealthCheckController(router)
	corsHandler := configs.GetCorsHandler()
	routerWithCors := corsHandler.Handler(router)
	return &CMSGoServer{
		serverConfig: serverConfig,
		router:       &routerWithCors,
	}
}
