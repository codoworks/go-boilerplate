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

func Post(c echo.Context) error {

	cc := c.(*context.Ctx) // custom context

	f := &models.CatForm{}

	if err := cc.BindAndValidate(f); err != nil {
		return cc.Err(err, nil)
	}

	// if err := c.Bind(f); err != nil {
	// 	return helpers.Error(c, constants.ERROR_BINDING_BODY, err)
	// }

	// if err := helpers.Validate(f); err != nil {
	// 	return c.JSON(http.StatusBadRequest, handlers.ValidationErrors(err))
	// }

	m := f.MapToModel()

	if err := m.Save(); err != nil {
		return cc.Err(err, nil)
	}

	return cc.Success(m.MapToForm())

}
