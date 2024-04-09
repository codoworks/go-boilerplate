/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package healthz

import (
	"github.com/codoworks/go-boilerplate/pkg/api/context"
	"github.com/codoworks/go-boilerplate/pkg/config"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {

	cc := c.(*context.Ctx) // custom context

	payload := map[string]string{
		"message": "ok",
		"version": config.Env.Version,
	}

	return cc.Success(payload)
}
