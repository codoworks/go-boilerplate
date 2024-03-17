/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print service version",
	Long:  `Print service version`,
	Run:   execVersionCmd,
}

func init() {
	// This is auto executed upon start
	// Initialization processes can go here ...

	// Register version command
	rootCmd.AddCommand(versionCmd)
}

func execVersionCmd(cmd *cobra.Command, args []string) {
	// Command execution goes here ...
	fmt.Printf("v%s\n", Version)
}
