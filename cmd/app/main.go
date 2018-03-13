package main

import (
	"errors"
	"net/http"

	redis "github.com/go-redis/redis"
	"github.com/go-zoo/bone"
	"github.com/sagarp-webonise/todo/app"
	"github.com/sagarp-webonise/todo/pkg/database"
	"github.com/sagarp-webonise/todo/pkg/logger"
	customRedis "github.com/sagarp-webonise/todo/pkg/redis"
	"github.com/sagarp-webonise/todo/pkg/session"
	"github.com/sagarp-webonise/todo/pkg/templates"
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

	//initialise redis client
	objectStorage := &customRedis.ObjectStorageWrapper{
		InitialiseRedisClient: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	}

	a := &app.App{
		Router:    router,
		Cfg:       cfg,
		Log:       log,
		TplParser: &templates.TemplateParser{},
		DB:        dbConn,
		Session:   &session.AppSession{},
		Redis:     objectStorage,
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
