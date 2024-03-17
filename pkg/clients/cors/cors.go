/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package cors

import (
	"reflect"

	"github.com/codoworks/go-boilerplate/pkg/config/features"
)

type CorsClient struct {
	name   string
	config features.CorsConfig
}

func (c *CorsClient) Name() string {
	return c.name
}

func (c *CorsClient) Configure(v any) {
	c.config = v.(reflect.Value).Interface().(features.CorsConfig)
}

func (c *CorsClient) GetConfig() features.CorsConfig {
	return c.config
}
