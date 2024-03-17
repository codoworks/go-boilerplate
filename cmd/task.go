/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package cmd

import (
	"github.com/codoworks/go-boilerplate/cmd/task"
	"github.com/codoworks/go-boilerplate/pkg/clients/logger"
	"github.com/codoworks/go-boilerplate/pkg/proc"

	"github.com/spf13/cobra"
)

// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use:   "task <option>",
	Short: "Start task",
	Long: `Run a one-time only task.
Please key in an option to start. Type 'task -h' for more information.

Popular options are:
- task list
- task init
- task cleanup

Use -d or --dev to start in development mode.`,
}

func init() {
	// This is auto executed upon start
	// Initialization processes can go here ...

	// Register sub commands
	taskCmd.AddCommand(task.ListCmd)
	taskCmd.AddCommand(task.ExecCmd)

	// Register persistent function for all sub commands
	taskCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		rootCmd.PersistentPreRun(cmd, args)
		execTaskPersistentPreRun()
	}

	// Register task command
	rootCmd.AddCommand(taskCmd)
}

func execTaskPersistentPreRun() {
	logger.Debug("Executing task persistent pre run ...")

	proc.InitClients()
	proc.ConfigureClients()
	proc.InitDbConnection()
	proc.InitModels()
	// You can initialize other features here ...
	// this will run before any command, make sure to put only global initializations here
	// to avoid running into nil pointers or undefined variables
	// ...
}
