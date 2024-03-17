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

// CreateCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create database",
	Long:  `Create dababase.`,
	Run:   execCreateCmd,
}

func init() {
	// This is auto executed upon start
	// Initialization processes can go here ...
}

func execCreateCmd(cmd *cobra.Command, args []string) {
	// Command execution goes here ...
	proc.DBCreate()
}
