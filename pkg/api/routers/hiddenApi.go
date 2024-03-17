/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package routers

import (
	"github.com/codoworks/go-boilerplate/pkg/api/handlers/errors"
	healthHandlers "github.com/codoworks/go-boilerplate/pkg/api/handlers/healthz"
	"github.com/codoworks/go-boilerplate/pkg/api/middlewares"
	"github.com/codoworks/go-boilerplate/pkg/clients/logger"
	"github.com/codoworks/go-boilerplate/pkg/config"
	"github.com/codoworks/go-boilerplate/pkg/utils/constants"
)

var hiddenApiRouter *Router

func InitHiddenAPIRouter() {
	logger.Debug("Initializing hidden api router ...")
	hiddenApiRouter = &Router{}
	hiddenApiRouter.Name = "hidden API"
	hiddenApiRouter.Init()

	// order is important here
	// first register development middlewares
	if config.DevModeFlag {
		logger.Debug("Registering hidden api development middlewares ...")
		registerHiddenApiDevModeMiddleware()
	}

	// next register middlwares
	logger.Debug("Registering hidden api middlewares ...")
	registerHiddenAPIMiddlewares()

	// next register all health check routes
	logger.Debug("Registering hidden api health routes ...")
	registerHiddenApiHealthCheckHandlers()

	// next register security related middleware
	logger.Debug("Registering hidden api security middlewares ...")
	registerHiddenApiSecurityMiddlewares()

	// next register all routes
	logger.Debug("Registering hidden api hidden routes ...")
	registerHiddenAPIRoutes()

	// finally register default fallback error handlers
	// 404 is handled here as the last route
	logger.Debug("Registering hidden api error handlers ...")
	registerHiddenApiErrorHandlers()

	logger.Debug("Hidden api registration complete.")
}

func HiddenAPIRouter() *Router {
	return hiddenApiRouter
}

func registerHiddenAPIMiddlewares() {
	hiddenApiRouter.RegisterPreMiddleware(middlewares.SlashesMiddleware())

	hiddenApiRouter.RegisterMiddleware(middlewares.LoggerMiddleware())
	hiddenApiRouter.RegisterMiddleware(middlewares.TimeoutMiddleware())
	hiddenApiRouter.RegisterMiddleware(middlewares.RequestHeadersMiddleware())
	hiddenApiRouter.RegisterMiddleware(middlewares.ResponseHeadersMiddleware())

	if config.Feature(constants.FEATURE_GZIP).IsEnabled() {
		hiddenApiRouter.RegisterMiddleware(middlewares.GzipMiddleware())
	}
}

func registerHiddenApiDevModeMiddleware() {
	hiddenApiRouter.RegisterMiddleware(middlewares.BodyDumpMiddleware())
}

func registerHiddenApiSecurityMiddlewares() {
	hiddenApiRouter.RegisterMiddleware(middlewares.XSSCheckMiddleware())

	if config.Feature(constants.FEATURE_CORS).IsEnabled() {
		hiddenApiRouter.RegisterMiddleware(middlewares.CORSMiddleware())
	}

	if config.Feature(constants.FEATURE_ORY_KRATOS).IsEnabled() {
		hiddenApiRouter.RegisterMiddleware(middlewares.AuthenticationMiddleware())
	}

	// if config.Feature(constants.FEATURE_ORY_KETO).IsEnabled() {
	// 	// keto middleware <- this will check if the user has the right permissions like system admin
	// 	// hiddenApiRouter.RegisterMiddleware(middlewares.AuthenticationMiddleware())
	// }
}

func registerHiddenApiErrorHandlers() {
	hiddenApiRouter.Echo.HTTPErrorHandler = errors.AutomatedHttpErrorHandler()
	hiddenApiRouter.Echo.RouteNotFound("/*", errors.NotFound)
}

func registerHiddenApiHealthCheckHandlers() {
	health := hiddenApiRouter.Echo.Group("/health")
	health.GET("/alive", healthHandlers.Index)
	health.GET("/ready", healthHandlers.Ready)
}

func registerHiddenAPIRoutes() {

	// add more routes here ...
}
