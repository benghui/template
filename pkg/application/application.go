package application

import (
	"github.com/template/pkg/config"
	"github.com/template/pkg/db"
	"github.com/template/pkg/logger"
	"github.com/template/pkg/router"
	"github.com/template/pkg/server"
)

// Application struct holds DB, server & configuration data for dependency injection
type Application struct {
	DB  *db.DB
	Cfg *config.Config
	Srv *server.Server
}

// GetApp captures env variables, initializes server & establishes DB connection then returns reference to all.
func GetApp() (*Application, error) {
	cfg := config.GetConfig()

	db, err := db.GetDB(cfg.GetDBConnStr())

	srv := server.
		GetServer().
		WithAddr(cfg.GetAPIPort()).
		WithRouter(router.GetRouter(db)).
		WithErrLogger(logger.Error)

	if err != nil {
		return nil, err
	}

	return &Application{
		DB: db,
		Cfg: cfg,
		Srv: srv,
	}, nil
}
