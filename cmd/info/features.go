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

// FeaturesCmd represents the features command
var FeaturesCmd = &cobra.Command{
	Use:   "features",
	Short: "Print service features configuration",
	Long:  `Print service features configuration based on environment`,
	Run:   execFeaturesCmd,
}

func init() {
	// This is auto executed upon start
	// Initialization processes can go here ...
}

func execFeaturesCmd(cmd *cobra.Command, args []string) {
	// Command execution goes here ...
	config.Env.PrintServiceFeatures()
}
