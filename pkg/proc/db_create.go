/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package proc

import (
	"fmt"

	"github.com/codoworks/go-boilerplate/pkg/clients/dbc"
	"github.com/codoworks/go-boilerplate/pkg/clients/logger"
)

func DBCreate() {
	// init feature [database]
	logger.SetLogger(string(logger.DebugLvl))

	dbClient := dbc.GetDBClient()

	dbClient.InitServerConnection()

	if err := dbClient.CreateDatabase(); err != nil {
		panic(fmt.Errorf("failed to create database: %w", err))
	}

	// logger.Info("Database '" + config.Env.Config.DBName + "' created successfully.")

}
