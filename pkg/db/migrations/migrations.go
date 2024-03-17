/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package migrations

import (
	"github.com/codoworks/go-boilerplate/pkg/clients/logger"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var Migrations *gormigrate.Gormigrate
var MigrationsList = []*gormigrate.Migration{}

func Init(db *gorm.DB) {
	Migrations = gormigrate.New(db, gormigrate.DefaultOptions, MigrationsList)
}

func AddMigration(migration *gormigrate.Migration) {
	MigrationsList = append(MigrationsList, migration)
}

func Migrate() error {
	return Migrations.Migrate()
}

func Rollback() error {
	return Migrations.RollbackLast()
}

func AutoMigrateAndLog(db *gorm.DB, model interface{}, id string) error {
	if err := db.AutoMigrate(model); err != nil {
		logFail(id, err)
		return err
	}
	logSuccess(id)
	return nil
}

func logSuccess(id string, rollback ...bool) {
	if len(rollback) > 0 && rollback[0] {
		logger.Info("Rolled back migration: %s", id)
		return
	}
	logger.Info("Applied migration: %s", id)
}

func logFail(id string, err error, rollback ...bool) {
	if len(rollback) > 0 && rollback[0] {
		logger.Error("Failed to rollback migration: %s, error: %s", id, err)
		return
	}
	logger.Error("Failed to apply migration: %s, error: %s", id, err)
}
