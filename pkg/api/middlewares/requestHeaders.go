/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package middlewares

import (
	"net/http"

	"github.com/codoworks/go-boilerplate/pkg/api/handlers"
	"github.com/codoworks/go-boilerplate/pkg/utils/constants"

	"github.com/labstack/echo/v4"
)

func RequestHeadersMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if (len(c.Request().Header["Accept"]) == 0) || c.Request().Header["Accept"][0] != "application/json" {
				r := handlers.BuildResponse(
					constants.STATUS_CODE_NOT_ACCEPTABLE_WITHOUT_ACCEPT_HEADER,
					constants.MSG_NOT_ACCEPTABLE,
					[]string{constants.MSG_MISSING_ACCEPT_HEADER},
					nil)
				return c.JSON(http.StatusNotAcceptable, r)
			}
			if c.Request().Method == "GET" {
				return next(c)
			}
			if (len(c.Request().Header["Content-Type"]) == 0) || c.Request().Header["Content-Type"][0] != "application/json" {
				r := handlers.BuildResponse(
					constants.STATUS_CODE_NOT_ACCEPTABLE_WITHOUT_CONTENT_TYPE_HEADER,
					constants.MSG_NOT_ACCEPTABLE,
					[]string{constants.MSG_MISSING_CONTENT_TYPE_HEADER},
					nil)
				return c.JSON(http.StatusNotAcceptable, r)
			}
			return next(c)
		}
	}
}
