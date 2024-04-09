/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package routers

import (
	catsHandlers "github.com/codoworks/go-boilerplate/pkg/api/handlers/cats"
	"github.com/codoworks/go-boilerplate/pkg/api/handlers/errors"
	healthHandlers "github.com/codoworks/go-boilerplate/pkg/api/handlers/healthz"
	usersHandlers "github.com/codoworks/go-boilerplate/pkg/api/handlers/users"
	"github.com/codoworks/go-boilerplate/pkg/api/middlewares"
	"github.com/codoworks/go-boilerplate/pkg/clients/logger"
	"github.com/codoworks/go-boilerplate/pkg/config"
	"github.com/codoworks/go-boilerplate/pkg/utils/constants"
)

var protectedApiRouter *Router

func InitProtectedAPIRouter() {
	logger.Debug("Initializing protected api router ...")
	protectedApiRouter = &Router{}
	protectedApiRouter.Name = "protected API"
	protectedApiRouter.Init()

	// first middleware to be registered must be the custom context middleware
	protectedApiRouter.RegisterMiddleware(middlewares.CustomContextMiddleware())

	// order is important here
	// first register development middlewares
	if config.DevModeFlag {
		logger.Debug("Registering protected api development middlewares ...")
		registerProtectedApiDevModeMiddleware()
	}

	// next register middlwares
	logger.Debug("Registering protected api middlewares ...")
	registerProtectedAPIMiddlewares()

	// next register all health check routes
	logger.Debug("Registering protected api health routes ...")
	registerProtectedApiHealthCheckHandlers()

	// next register security related middleware
	logger.Debug("Registering protected api security middlewares ...")
	registerProtectedApiSecurityMiddlewares()

	// next register all routes
	logger.Debug("Registering protected api protected routes ...")
	registerProtectedAPIRoutes()

	// finally register default fallback error handlers
	// 404 is handled here as the last route
	logger.Debug("Registering protected api error handlers ...")
	registerProtectedApiErrorHandlers()

	logger.Debug("Protected api registration complete.")
}

func ProtectedAPIRouter() *Router {
	return protectedApiRouter
}

func registerProtectedAPIMiddlewares() {
	protectedApiRouter.RegisterPreMiddleware(middlewares.SlashesMiddleware())

	protectedApiRouter.RegisterMiddleware(middlewares.LoggerMiddleware())
	protectedApiRouter.RegisterMiddleware(middlewares.TimeoutMiddleware())
	protectedApiRouter.RegisterMiddleware(middlewares.RequestHeadersMiddleware())
	protectedApiRouter.RegisterMiddleware(middlewares.ResponseHeadersMiddleware())

	if config.Feature(constants.FEATURE_GZIP).IsEnabled() {
		protectedApiRouter.RegisterMiddleware(middlewares.GzipMiddleware())
	}
}

func registerProtectedApiDevModeMiddleware() {
	protectedApiRouter.RegisterMiddleware(middlewares.BodyDumpMiddleware())
}

func registerProtectedApiSecurityMiddlewares() {
	protectedApiRouter.RegisterMiddleware(middlewares.XSSCheckMiddleware())

	if config.Feature(constants.FEATURE_CORS).IsEnabled() {
		protectedApiRouter.RegisterMiddleware(middlewares.CORSMiddleware())
	}

	if config.Feature(constants.FEATURE_ORY_KRATOS).IsEnabled() {
		protectedApiRouter.RegisterMiddleware(middlewares.AuthenticationMiddleware())
	}

	if config.Feature(constants.FEATURE_ORY_KETO).IsEnabled() {
		// keto middleware <- this will check if the user has the right permissions like system admin
		// protectedApiRouter.RegisterMiddleware(middlewares.AuthenticationMiddleware())
	}
}

func registerProtectedApiErrorHandlers() {
	protectedApiRouter.Echo.HTTPErrorHandler = errors.AutomatedHttpErrorHandler()
	protectedApiRouter.Echo.RouteNotFound("/*", errors.NotFound)
}

func registerProtectedApiHealthCheckHandlers() {
	health := protectedApiRouter.Echo.Group("/health")
	health.GET("/alive", healthHandlers.Index)
	health.GET("/ready", healthHandlers.Ready)
}

func registerProtectedAPIRoutes() {
	cats := protectedApiRouter.Echo.Group("/cats")
	cats.GET("", catsHandlers.Index)
	cats.GET("/:id", catsHandlers.Get)
	cats.POST("", catsHandlers.Post)
	cats.PUT("/:id", catsHandlers.Put)
	cats.DELETE("/:id", catsHandlers.Delete)

	users := protectedApiRouter.Echo.Group("/users")
	users.GET("", usersHandlers.Index)
	users.GET("/:id", usersHandlers.Get)
	users.POST("", usersHandlers.Post)
	// users.PUT("/:id", usersHandlers.Put)
	users.DELETE("/:id", usersHandlers.Delete)

	// add more routes here ...
}
