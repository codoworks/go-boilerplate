/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package cmd

import (
	"github.com/codoworks/go-boilerplate/cmd/start"
	"github.com/codoworks/go-boilerplate/pkg/clients/logger"
	"github.com/codoworks/go-boilerplate/pkg/config"
	"github.com/codoworks/go-boilerplate/pkg/proc"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start <option>",
	Short: "Start service or daemon",
	Long: `Start web server, daemons or other services.

When running this command without options, it will start all http servers.
If you wish to start a single server instead, you can choose from the 
available options.

Use -d or --dev to start in development mode.`,

	// Uncomment the following line if you wish to execute the command
	// by default this command does not execute anything as it is a parent command
	//
	Run: execStartCmd,
}

func init() {
	// This is auto executed upon start
	// Initialization processes can go here ...

	// Register sub commands
	startCmd.AddCommand(start.PublicApiCmd)
	startCmd.AddCommand(start.ProtectedApiCmd)
	startCmd.AddCommand(start.HiddenApiCmd)
	startCmd.AddCommand(start.WatcherCmd)

	// Set global flags
	startCmd.PersistentFlags().BoolVar(&config.StartWatcherFlag, "watcher", false, "Start watcher daemon in background")
	startCmd.PersistentFlags().StringVarP(&config.HostFlag, "host", "H", "", "Service host")
	startCmd.PersistentFlags().StringVar(&config.ProtectedPortFlag, "protected-api-port", "", "Protected API Service port")
	startCmd.PersistentFlags().StringVar(&config.PublicPortFlag, "public-api-port", "", "Public API Service port")
	startCmd.PersistentFlags().StringVar(&config.HiddenPortFlag, "hidden-api-port", "", "Hidden API Service port")

	// Register persistent function for all sub commands
	startCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		rootCmd.PersistentPreRun(cmd, args)
		execStartPersistentPreRun()
	}

	// Register start command
	rootCmd.AddCommand(startCmd)
}

func execStartPersistentPreRun() {
	logger.Debug("Executing start persistent pre run ...")

	proc.InitClients()
	proc.ConfigureClients()
	proc.InitDbConnection()
	proc.InitModels()
	// You can initialize other features here ...
	// this will run before any command, make sure to put only global initializations here
	// to avoid running into nil pointers or undefined variables
	// ...
}

func execStartCmd(cmd *cobra.Command, args []string) {
	// Command execution goes here ...
	if config.StartWatcherFlag {
		go start.WatcherCmd.Run(cmd, args)
	}
	cmd.Flags().Set("watcher", "false")
	go start.HiddenApiCmd.Run(cmd, args)
	go start.PublicApiCmd.Run(cmd, args)
	start.ProtectedApiCmd.Run(cmd, args)
}
