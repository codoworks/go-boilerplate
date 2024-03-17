/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package features

import "github.com/codoworks/go-boilerplate/pkg/utils/constants"

type KratosConfig struct {
	PublicService string `mapstructure:"KRATOS_PUBLIC_SERVICE"`
	AdminService  string `mapstructure:"KRATOS_ADMIN_SERVICE"`
}

var kratos = &Feature{
	Name:       constants.FEATURE_ORY_KRATOS,
	Config:     &KratosConfig{},
	enabled:    true,
	configured: false,
	ready:      false,
	requirements: []string{
		"PublicService",
		"AdminService",
	},
}

func init() {
	Features.Add(kratos)
}
