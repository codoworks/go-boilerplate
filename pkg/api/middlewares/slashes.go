/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SlashesMiddleware() echo.MiddlewareFunc {
	return middleware.RemoveTrailingSlash()
}
