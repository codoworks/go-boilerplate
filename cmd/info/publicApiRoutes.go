/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package info

import (
	"github.com/codoworks/go-boilerplate/pkg/proc"

	"github.com/spf13/cobra"
)

// PublicApiRoutesCmd represents the public-api-routes command
var PublicApiRoutesCmd = &cobra.Command{
	Use:   "public-api-routes",
	Short: "Print public API routes",
	Long:  `Print all public API routes table`,
	Run:   execPublicApiRoutesCmd,
}

func init() {
	// This is auto executed upon start
	// Initialization processes can go here ...
}

func execPublicApiRoutesCmd(cmd *cobra.Command, args []string) {
	// Command execution goes here ...
	proc.PrintPublicRoutesTable()
}
