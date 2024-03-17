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

// RollbackCmd represents the rollback command
var RollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback database",
	Long:  `Rollback one database migration.`,
	Run:   execRollbackCmd,
}

func init() {
	// This is auto executed upon start
	// Initialization processes can go here ...
}

func execRollbackCmd(cmd *cobra.Command, args []string) {
	// Command execution goes here ...
	proc.DBRollback()
}
