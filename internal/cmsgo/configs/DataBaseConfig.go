package configs

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"CMSGo-backend/internal/cmsgo/env"
)

type DBConnectionInfo struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	DbName   string `toml:"db_name"`
}

type DataBaseConfig struct {
	Database DBConnectionInfo
}

func GetDataBaseConnection() *gorm.DB {
	dbConfig := getDataBaseConfig()
	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Database.Host, dbConfig.Database.Port, dbConfig.Database.Username, dbConfig.Database.Password, dbConfig.Database.DbName)
	dbConnection, err := gorm.Open(postgres.Open(dbURI))
	if err != nil {
		log.Fatal("Couldn't connect to DB: ", err)
	} else {
		log.Println("Successfully connected to DB")
	}
	return dbConnection
}

func getDataBaseConfig() *DataBaseConfig {
	config := DataBaseConfig{}
	if _, err := toml.DecodeFile(env.Find("config-path"), &config); err != nil {
		log.Fatal("Couldn't read configs for database\n")
	}
	return &config
}
