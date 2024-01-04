package main

import (
	"github.com/F1zm0n/bookey/internal/bookey/delivery/http/handlers"
	"github.com/F1zm0n/bookey/internal/bookey/repository/database/connect"
	"github.com/F1zm0n/bookey/internal/bookey/repository/migrations/sqlc"
	"github.com/F1zm0n/bookey/pkg/config"
	logging "github.com/F1zm0n/bookey/pkg/logger"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
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

	dbtx := sqlc.New(db)

	logger.Info("creating router")
	router := gin.Default()

	handler := handlers.NewHandler(logger, db, dbtx)
	handler.Register(router)

	router.Run(":" + cfg.Server.Port)
}
