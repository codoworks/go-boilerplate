/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package healthz

import (
	"github.com/codoworks/go-boilerplate/pkg/api/context"
	"github.com/codoworks/go-boilerplate/pkg/clients/dbc"

	"github.com/labstack/echo/v4"
)

func Ready(c echo.Context) error {

	cc := c.(*context.Ctx) // custom context

	dbClient := dbc.GetDBClient()
	if err := dbClient.Ping(); err != nil {
		return cc.Err(err, nil)
	}

	payload := map[string]string{
		"message": "ready",
	}
	return cc.Success(payload)
}
