package connect

import (
	logging "bookey/pkg/logger"
	"database/sql"
	"fmt"
)

func Connection(driver, user, name, password, sslmode string, Logger *logging.Logger) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", user, name, password, sslmode)

	Logger.Infof("starting database connection, %v", connStr)

	db, err := sql.Open(driver, connStr)
	if err != nil {
		Logger.Errorf("error connecting database, %v", connStr)
		return db, err
	}
	return db, nil
}
