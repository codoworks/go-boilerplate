/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package middlewares

import (
	"github.com/codoworks/go-boilerplate/pkg/utils/constants"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoggerMiddleware() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom} ${user_agent} ${remote_ip} ${status} ${method} ${uri} ${latency_human} ${bytes_out}\n",
		CustomTimeFormat: constants.DEFAULT_LOGGER_TIMESTAMP_FORMAT,
	})
}

// Available options
//
// - time_unix
// - time_unix_milli
// - time_unix_micro
// - time_unix_nano
// - time_rfc3339
// - time_rfc3339_nano
// - time_custom
// - id (Request ID)
// - remote_ip
// - uri
// - host
// - method
// - path
// - protocol
// - referer
// - user_agent
// - status
// - error
// - latency (In nanoseconds)
// - latency_human (Human readable)
// - bytes_in (Bytes received)
// - bytes_out (Bytes sent)
// - header:<NAME>
// - query:<NAME>
// - form:<NAME>
