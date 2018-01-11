package app

import (
	"database/sql"

	"github.com/go-zoo/bone"
	"github.com/kaddiya/todo/pkg/logger"
	"github.com/kaddiya/todo/pkg/templates"
)

// App enscapsulates the App environment
type App struct {
	Router    *bone.Mux
	Cfg       *Config
	Log       logger.ILogger
	TplParser templates.ITemplateParser
	DB        *sql.DB
}
