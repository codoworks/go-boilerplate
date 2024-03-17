/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package db

import (
	"fmt"

	"github.com/codoworks/go-boilerplate/pkg/config"
	"github.com/codoworks/go-boilerplate/pkg/proc"

	"github.com/spf13/cobra"
)

var DropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop database",
	Long:  `Drop dababase.`,
	Run:   execDropCmd,
}

func init() {
	// This is auto executed upon start
	// Initialization processes can go here ...
	DropCmd.PersistentFlags().BoolVar(&config.ConfirmFlag, "confirm", false, "Confirm deletion of database")
}

func execDropCmd(cmd *cobra.Command, args []string) {
	// Command execution goes here ...
	if !config.ConfirmFlag {
		fmt.Println("This is a destructive action and it is irreversible.")
		fmt.Printf("Are you sure you want to drop the database? \n")

		fmt.Printf("To delete please run again using the `--confirm` flag \n")
		return
	}
	proc.DBDrop()
}
