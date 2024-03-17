/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Post(c echo.Context) error {
	// newUser := &models.UserObj{}
	// if err := c.Bind(newUser); err != nil {
	// 	c.Echo().Logger.Error(constants.ERROR_BINDING_BODY)
	// 	c.Echo().Logger.Error(err)
	// 	return constants.ERROR_BINDING_BODY
	// }
	// if err := Validate.Struct(newUser); err != nil {
	// 	r := handlers.BuildValidationErrorsResponse(err.(validator.ValidationErrors))
	// 	return c.JSON(http.StatusBadRequest, r)
	// }
	// identity, err := kratos.Cli.CreateIdentity(newUser.ConvertToMap(), newUser.Password)
	// if err != nil {
	// 	return err
	// }
	// return c.JSON(http.StatusOK, handlers.Success(identity))
	return c.JSON(http.StatusOK, nil)
}
