/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package dbc

import (
	"github.com/codoworks/go-boilerplate/pkg/clients/dbc/adapters"
	"github.com/codoworks/go-boilerplate/pkg/utils/constants"

	"gorm.io/gorm"
	gLogger "gorm.io/gorm/logger"
)

var dbClient *DBClient

func init() {
	dbClient = &DBClient{
		name:    constants.FEATURE_DATABASE,
		adapter: adapters.Adapters,
		silent:  true,
		gormConfig: &gorm.Config{
			Logger: gLogger.Default.LogMode(gLogger.Silent),
		},
	}
}

func GetDBClient() *DBClient {
	return dbClient
}
