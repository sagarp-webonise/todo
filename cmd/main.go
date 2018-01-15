package main

import (
	"errors"
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/kaddiya/todo/app"
	"github.com/kaddiya/todo/app/exp"
	"github.com/kaddiya/todo/pkg/database"
	"github.com/kaddiya/todo/pkg/logger"
	"github.com/kaddiya/todo/pkg/templates"
)

func main() {

	//initialise the router
	router := bone.New()

	//initialise logger
	log := &logger.RealLogger{}
	log.Initialise()

	// need to constrcut an instance of the AppConfig from various environment vars
	cfg, cfgErr := constructAppConfig()
	//hydrate the map of DB connection params and send it
	dbConnectionParams := make(map[string]string)
	db := &database.DatabaseWrapper{}

	dbConn, dbErr := db.Initialise(dbConnectionParams)
	if dbErr != nil || dbConn == nil {
		panic(errors.New("could not initialise the DB"))
	}

	//if the configuration is not loaded then exit before startup
	if cfgErr != nil {
		panic("the configuration wasnt enabled")
	}

	a := &app.App{
		Router:         router,
		Cfg:            cfg,
		Log:            log,
		TplParser:      &templates.TemplateParser{},
		DB:             dbConn,
		TodoSeviceImpl: &exp.TodoServiceImpl{DB: dbConn},
	}

	a.InitRouter()
	err := http.ListenAndServe(cfg.Port, router)
	if err != nil {
		panic(err)
	}
}

func constructAppConfig() (*app.Config, error) {
	cfg := &app.Config{
		Port: ":9999",
	}
	return cfg, nil
}
