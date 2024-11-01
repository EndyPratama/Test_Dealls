package main

import (
	"test_dealls/src/business/domain"
	"test_dealls/src/business/usecase"
	"test_dealls/src/handler"
	"test_dealls/src/utils/config"
	"test_dealls/src/utils/configreader"
	"test_dealls/src/utils/log"
	"test_dealls/src/utils/sql"
)

const (
	configfile   string = "./etc/cfg/conf.json"
	appnamespace string = "test_dealls"
)

func main() {
	// init config
	cfg := config.Init()
	configreader := configreader.Init(configreader.Options{
		Name: "conf",
		Type: "yaml",
		Path: "./etc/cfg",
	})
	configreader.ReadConfig(&cfg)

	// init logger
	log := log.Init(cfg.Log)

	// init db conn
	db := sql.Init(cfg.SQL, log)

	// init all domain
	d := domain.Init(log, db, cfg)

	// init all uc
	uc := usecase.Init(log, d, cfg)

	r := handler.Init(cfg, configreader, log, uc)

	r.Run()
}
