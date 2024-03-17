/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package users

import (
	"net/http"

	"github.com/codoworks/go-boilerplate/pkg/api/handlers"
	"github.com/codoworks/go-boilerplate/pkg/clients/kratos"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	kratosCli := kratos.GetClient()
	identities, err := kratosCli.GetAllIdentity()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, handlers.Success(identities))
}

// Delete handler
