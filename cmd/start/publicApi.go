/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package start

import (
	"github.com/codoworks/go-boilerplate/pkg/config"
	"github.com/codoworks/go-boilerplate/pkg/proc"

	"github.com/spf13/cobra"
)

// PublicApiCmd represents the publicApi command
var PublicApiCmd = &cobra.Command{
	Use:   "publicApi",
	Short: "Start public API service",
	Long:  `Start public API web server.`,
	Run:   execPublicApiCmd,
}

func init() {
	// This is auto executed upon start
	// Initialization processes can go here ...
}

func execPublicApiCmd(cmd *cobra.Command, args []string) {
	// Command execution goes here ...
	if config.StartWatcherFlag {
		go WatcherCmd.Run(cmd, args)
	}
	proc.StartPublicApi()
}
