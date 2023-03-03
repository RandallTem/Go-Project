package configs

import (
	"github.com/rs/cors"
	"log"

	"github.com/BurntSushi/toml"

	"CMSGo-backend/internal/cmsgo/env"
)

type ServerConfig struct {
	ServerPort string `toml:"server_port"`
}

type CorsConfig struct {
	AllowedSources []string `toml:"allowed_sources"`
}

func GetServerConfig() *ServerConfig {
	serverConfig := ServerConfig{
		ServerPort: ":8080",
	}
	if _, err := toml.DecodeFile(env.Find("config-path"), &serverConfig); err != nil {
		log.Printf("Couldn't read server configs. Used default value %s\n", serverConfig.ServerPort)
	}
	return &serverConfig
}

func GetCorsHandler() *cors.Cors {
	corsConfig := CorsConfig{
		AllowedSources: []string{"*"},
	}
	if _, err := toml.DecodeFile(env.Find("config-path"), &corsConfig); err != nil {
		log.Printf("Couldn't read cors configs. All sources are allowed\n")
	}
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: corsConfig.AllowedSources,
		AllowedMethods: []string{"GET", "POST"},
	})
	return corsHandler
}
