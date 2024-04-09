/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package cats

import (
	"github.com/codoworks/go-boilerplate/pkg/api/context"
	"github.com/codoworks/go-boilerplate/pkg/db/models"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {

	cc := c.(*context.Ctx) // custom context

	id := c.Param("id")

	if ok := cc.ValidateID(id); !ok {
		return cc.IdIsInvalid()
	}

	m, err := models.CatModel().Find(id)

	if err != nil {
		return cc.Err(err, nil)
	}

	return cc.Success(m.MapToForm())

}
