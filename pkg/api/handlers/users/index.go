/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package users

import (
	"github.com/codoworks/go-boilerplate/pkg/api/context"
	"github.com/codoworks/go-boilerplate/pkg/clients/kratos"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {

	cc := c.(*context.Ctx) // custom context

	kratosCli := kratos.GetClient()

	identities, err := kratosCli.GetAllIdentity()

	if err != nil {
		return cc.Err(err, nil)
	}

	return cc.Success(identities)

}

// Delete handler
