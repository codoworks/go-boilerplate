/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package start

import (
	"github.com/codoworks/go-boilerplate/pkg/proc"

	"github.com/spf13/cobra"
)

// WatcherCmd represents the watcher command
var WatcherCmd = &cobra.Command{
	Use:   "watcher",
	Short: "Start watcher daemon",
	Long:  `Start watcher daemon.`,
	Run:   execWatcherCmd,
}

func init() {
	// This is auto executed upon start
	// Initialization processes can go here ...
}

func execWatcherCmd(cmd *cobra.Command, args []string) {
	// Command execution goes here ...
	proc.StartWatcher()
}
