/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package helpers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var googleValidate *validator.Validate

func init() {
	googleValidate = validator.New()
}

func Validate(s interface{}) error {
	return googleValidate.Struct(s)
}

func Error(c echo.Context, err error, ori_err error) error {
	c.Logger().Error(err, ori_err)
	return err
}
