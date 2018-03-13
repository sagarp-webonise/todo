package app

import (
	"database/sql"

	redisClient "github.com/go-redis/redis"
	"github.com/go-zoo/bone"
	"github.com/sagarp-webonise/todo/pkg/logger"
	redis "github.com/sagarp-webonise/todo/pkg/redis"
	"github.com/sagarp-webonise/todo/pkg/session"
	"github.com/sagarp-webonise/todo/pkg/templates"
)

// App enscapsulates the App environment
type App struct {
	Router      *bone.Mux
	Cfg         *Config
	Log         logger.ILogger
	TplParser   templates.ITemplateParser
	DB          *sql.DB
	Session     session.Session
	Redis       redis.IRedis
	RedisClient *redisClient.Client
}
