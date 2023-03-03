package main

import (
	"CMSGo-backend/internal/cmsgo"
	"flag"
	"log"

	"CMSGo-backend/internal/cmsgo/env"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/application_config.toml", "Config path")
}

func configureEnvironment() {
	flag.Parse()
	env.Add("config-path", configPath)
}

func main() {
	configureEnvironment()
	if err := cmsgo.StartServer(); err != nil {
		log.Fatal("Couldn't start server: ", err)
	}
}
