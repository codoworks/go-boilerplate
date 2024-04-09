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

func Get(c echo.Context) error {

	cc := c.(*context.Ctx) // custom context

	id := c.Param("id")

	if ok := cc.ValidateID(id); !ok {
		return cc.IdIsInvalid()
	}

	kratosCli := kratos.GetClient()

	identity, err := kratosCli.GetIdentity(id)

	if err != nil {
		return cc.Err(err, nil)
	}

	return cc.Success(identity)

}
