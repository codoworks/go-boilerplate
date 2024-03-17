/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package features

import "github.com/codoworks/go-boilerplate/pkg/utils/constants"

type KetoConfig struct {
	ReadService  string `mapstructure:"KETO_READ_SERVICE"`
	WriteService string `mapstructure:"KETO_WRITE_SERVICE"`
}

var keto = &Feature{
	Name:       constants.FEATURE_ORY_KETO,
	Config:     &KetoConfig{},
	enabled:    true,
	configured: false,
	ready:      false,
	requirements: []string{
		"ReadService",
		"WriteService",
	},
}

func init() {
	Features.Add(keto)
}
