package main

import (
	"fmt"
	"github.com/fwidjaya20/ecommerce-point-system/config"
	"github.com/fwidjaya20/ecommerce-point-system/internal/globals"
	"github.com/fwidjaya20/ecommerce-point-system/lib/database"
	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"
	"os"
)

func main() {
	var logger log.Logger

	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestamp)
	logger = log.With(logger, "caller", log.DefaultCaller)

	con := globals.DB()
	defer con.Close()

	initMigration(con)
}

func initMigration(dbConn *sqlx.DB) {
	root, err := os.Getwd()
	if nil != err {
		panic(fmt.Sprintf("failed retrieve root path : %v", err.Error()))
	}

	migrationPath := fmt.Sprintf("%s/%s", root, config.GetEnv(config.MIGRATION_PATH))
	database.Migrate(dbConn.DB, config.GetEnv(config.DB_NAME), migrationPath)
}
