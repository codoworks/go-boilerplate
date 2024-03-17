/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package proc

import (
	"github.com/codoworks/go-boilerplate/pkg/clients/dbc"
	"github.com/codoworks/go-boilerplate/pkg/clients/logger"
	"github.com/codoworks/go-boilerplate/pkg/db/migrations"
)

func DBMigrate() {
	logger.SetLogger(string(logger.DebugLvl))

	dbClient := dbc.GetDBClient()

	dbClient.InitDBConnection()

	migrations.Init(dbClient.DB)

	if err := migrations.Migrate(); err != nil {
		logger.Error("Failed to apply migrations: %s", err)
	}

	logger.Info("Migrations applied successfully")

}
