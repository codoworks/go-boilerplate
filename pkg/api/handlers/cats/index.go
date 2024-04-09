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

func Index(c echo.Context) error {

	cc := c.(*context.Ctx) // custom context

	ms, err := models.CatModel().FindAll()

	if err != nil {
		return cc.Err(err, nil)
	}

	var payload []*models.CatForm

	for _, m := range ms {
		payload = append(payload, m.MapToForm())
	}

	return cc.Success(payload)

}
