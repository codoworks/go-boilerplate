/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package info

import (
	"github.com/codoworks/go-boilerplate/pkg/config"

	"github.com/spf13/cobra"
)

// EnvCmd represents the env command
var EnvCmd = &cobra.Command{
	Use:   "env",
	Short: "Print service env",
	Long:  `Print service env`,
	Run:   execEnvCmd,
}

func init() {
	// This is auto executed upon start
	// Initialization processes can go here ...
}

func execEnvCmd(cmd *cobra.Command, args []string) {
	// Command execution goes here ...
	config.Env.PrintEnvironment()
}
