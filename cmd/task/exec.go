/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package task

import (
	"github.com/codoworks/go-boilerplate/pkg/proc"

	"github.com/spf13/cobra"
)

// ExecCmd represents the exec command
var ExecCmd = &cobra.Command{
	Use:   "exec",
	Short: "Start exec task",
	Long:  `Start the exec task.`,
	Run:   execExecCmd,
}

func init() {
	// This is auto executed upon start
	// Initialization processes can go here ...
}

func execExecCmd(cmd *cobra.Command, args []string) {
	// Command execution goes here ...
	proc.TaskExec(args)
}
