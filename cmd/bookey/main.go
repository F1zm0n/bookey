package main

import (
	"bookey/internal/bookey/repository/database/connect"
	"bookey/pkg/config"
	logging "bookey/pkg/logger"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx"
)

func main() {
	logger := logging.GetLogger()

	cfg := config.GetConfig()
	PostgresCfg := cfg.Database.Postgres

	db, err := connect.Connection(PostgresCfg.Driver, PostgresCfg.User, PostgresCfg.Name, PostgresCfg.Password, PostgresCfg.Sslmode, logger)
	if err != nil {
		logger.Errorf("error connecting to database %v", err)
		return
	}
	defer db.Close()

	logger.Info("creating router")
	router := gin.Default()

	router.Run(cfg.Server.Port)
}
