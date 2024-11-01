package config

import (
	"test_dealls/src/utils/log"
	"test_dealls/src/utils/sql"
	"time"
)

type Application struct {
	Gin  GinConfig
	SQL  sql.Config
	Meta ApplicationMeta
	Log  log.Config
}

func Init() Application {
	return Application{}
}

type GinConfig struct {
	Host    string
	Port    string
	Mode    string
	Timeout time.Duration
}

type ApplicationMeta struct {
	Title       string
	Description string
	Host        string
	BasePath    string
	Version     string
}
