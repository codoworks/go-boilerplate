/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package routers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/codoworks/go-boilerplate/pkg/clients/logger"
	"github.com/codoworks/go-boilerplate/pkg/utils/constants"

	"github.com/labstack/echo/v4"
)

type Router struct {
	// echo.Echo // change how echo is embedded
	Echo *echo.Echo
	Name string
}

func (r *Router) Init() {
	r.Echo = echo.New()
	r.Echo.HideBanner = true
	r.Echo.Logger = logger.GetLogger()
}

func (r *Router) RegisterPreMiddleware(middleware echo.MiddlewareFunc) {
	r.Echo.Pre(middleware)
}

func (r *Router) RegisterMiddleware(middleware echo.MiddlewareFunc) {
	r.Echo.Use(middleware)
}

func (r *Router) Start(host string, port string) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		r.Echo.Logger.Info(fmt.Sprintf("Starting %s server on port: %s", r.Name, port))
		if err := r.Echo.Start(host + ":" + port); err != nil && err != http.ErrServerClosed {
			r.Echo.Logger.Fatal(err)
			r.Echo.Logger.Fatal(constants.MSG_SERVER_SHUTTING_DOWN)
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if err := r.Echo.Shutdown(ctx); err != nil {
		r.Echo.Logger.Fatal(err)
	}
}
