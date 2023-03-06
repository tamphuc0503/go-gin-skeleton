package common

import (
	"os"
	"strings"
	"sync"
)

var lock = &sync.Mutex{}

type Settings struct {
	DbType        string
	DbName        string
	DbUser        string
	DbHost        string
	DbPassword    string
	DbPort        string
	IsDevelopment bool
}

var singleton *Settings

func CurrentSettings() *Settings {
	if singleton == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleton == nil {
			singleton = initializeSettings()
		}
	}
	return singleton
}

func initializeSettings() *Settings {
	var settings = &Settings{}
	settings.IsDevelopment = isDevelopment()
	settings.DbName = os.Getenv("DB_NAME")
	settings.DbType = os.Getenv("DB_TYPE")
	settings.DbPassword = os.Getenv("DB_PASSWORD")
	settings.DbPort = os.Getenv("DB_PORT")
	settings.DbHost = os.Getenv("DB_HOST")
	return settings
}

func isDevelopment() bool {
	return strings.ToLower(os.Getenv("APP_ENV")) == "development" || strings.ToLower(os.Getenv("GO_ENV")) == "development"
}
