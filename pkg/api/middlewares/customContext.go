/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package middlewares

import (
	"github.com/codoworks/go-boilerplate/pkg/api/context"
	"github.com/labstack/echo/v4"
)

// custom middleware template
func CustomContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &context.Ctx{
				Context: c,
			}
			return next(cc)
		}
	}
}
