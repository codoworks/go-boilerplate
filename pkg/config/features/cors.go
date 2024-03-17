/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package features

import "github.com/codoworks/go-boilerplate/pkg/utils/constants"

type CorsConfig struct {
	AllowOrigins  string `mapstructure:"CORS_ALLOW_ORIGINS"`
	AllowMethods  string `mapstructure:"CORS_ALLOW_METHODS"`
	AllowHeaders  string `mapstructure:"CORS_ALLOW_HEADERS"`
	MaxAge        int    `mapstructure:"CORS_MAX_AGE"`
	AllowCreds    bool   `mapstructure:"CORS_ALLOW_CREDENTIALS"`
	ExposeHeaders string `mapstructure:"CORS_EXPOSE_HEADERS"`
}

var cors = &Feature{
	Name:       constants.FEATURE_CORS,
	Config:     &CorsConfig{},
	enabled:    true,
	configured: false,
	ready:      false,
	requirements: []string{
		"AllowOrigins",
	},
}

func init() {
	Features.Add(cors)
}
