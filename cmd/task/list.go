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

// ListCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Long:  `List all available tasks.`,
	Run:   execListCmd,
}

func init() {
	// This is auto executed upon start
	// Initialization processes can go here ...
}

func execListCmd(cmd *cobra.Command, args []string) {
	// Command execution goes here ...
	proc.TaskList()
}
