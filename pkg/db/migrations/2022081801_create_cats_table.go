/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package migrations

import (
	"github.com/codoworks/go-boilerplate/pkg/db/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	m := &gormigrate.Migration{}

	m.ID = "2022081801_create_cats_table"

	m.Migrate = func(db *gorm.DB) error {
		type Cat struct {
			models.ModelBase
			Name string `gorm:"size:255"`
			Type string `gorm:"size:255"`
		}

		return AutoMigrateAndLog(db, &Cat{}, m.ID)
	}

	m.Rollback = func(db *gorm.DB) error {
		if err := db.Migrator().DropTable("cats"); err != nil {
			logFail(m.ID, err, true)
		}
		logSuccess(m.ID, true)
		return nil
	}

	AddMigration(m)
}
