/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package service

import (
	"reflect"

	"github.com/codoworks/go-boilerplate/pkg/config/features"
)

type ServiceClient struct {
	name   string
	config features.ServiceConfig
}

func (c *ServiceClient) Name() string {
	return c.name
}

func (c *ServiceClient) Configure(v any) {
	c.config = v.(reflect.Value).Interface().(features.ServiceConfig)
}

func (c *ServiceClient) GetConfig() features.ServiceConfig {
	return c.config
}
