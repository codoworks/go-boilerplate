/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package middlewares

import (
	"fmt"

	"github.com/codoworks/go-boilerplate/pkg/clients/kratos"
	"github.com/codoworks/go-boilerplate/pkg/utils/constants"

	"github.com/labstack/echo/v4"
)

func AuthenticationMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// skip authentication for health check
			if c.Path() == constants.NAME_HEALTH_PATH || c.Path() == fmt.Sprintf("%s%s", constants.NAME_HEALTH_PATH, constants.NAME_HEALTH_READY_PATH) {
				return next(c)
			}
			// validate session
			kratosCli := kratos.GetClient()
			session, err := kratosCli.ValidateSession(c.Request())
			if err != nil {
				c.Logger().Warn(err)
				c.Logger().Error(constants.ERROR_SESSION_NOT_FOUND)
				return constants.ERROR_NOT_AUTHORIZED
			}
			if !*session.Active {
				return constants.ERROR_NOT_AUTHORIZED
			}
			c.Logger().Warn("Session found:")
			c.Logger().Warn(session)
			kratosCli.Session.SetSession(session)
			return next(c)
		}
	}
}
