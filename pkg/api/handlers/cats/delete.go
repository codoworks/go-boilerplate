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

func Delete(c echo.Context) error {

	cc := c.(*context.Ctx) // custom context

	id := cc.Param("id")

	if ok := cc.ValidateID(id); !ok {
		return cc.IdIsInvalid()
	}

	if err := models.CatModel().Delete(id); err != nil {
		return cc.Err(err, nil)
	}

	return cc.Deleted()

}
