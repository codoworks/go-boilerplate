/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package db

import (
	"github.com/codoworks/go-boilerplate/pkg/proc"

	"github.com/spf13/cobra"
)

// MigrateCmd represents the migrate command
var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate database",
	Long:  `Run database migrations.`,
	Run:   execMigrateCmd,
}

func init() {
	// This is auto executed upon start
	// Initialization processes can go here ...
}

func execMigrateCmd(cmd *cobra.Command, args []string) {
	// Command execution goes here ...
	proc.DBMigrate()
}
