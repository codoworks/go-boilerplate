/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package context

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/codoworks/go-boilerplate/pkg/config"
	"github.com/codoworks/go-boilerplate/pkg/utils/constants"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Ctx struct {
	echo.Context
	CurrentUserID string
}

func (c *Ctx) ValidateID(id string) (ok bool) {
	// todo:- check the id format
	return len(id) == 36
}

func (c *Ctx) Err(err, originalErr error) error {
	c.Logger().Error(err, originalErr)
	return err
}

func (c *Ctx) IdIsInvalid() error {
	return c.Err(constants.ERROR_ID_NOT_FOUND, nil)
}

func (c *Ctx) Success(payload interface{}) error {
	return c.JSON(
		http.StatusOK,
		c.BuildResponse(
			constants.STATUS_CODE_SERVICE_SUCCESS,
			constants.MSG_SUCCESS,
			[]string{},
			payload,
		),
	)
}

func (c *Ctx) Accepted() error {
	return c.JSON(
		http.StatusAccepted,
		c.BuildResponse(
			constants.STATUS_CODE_SERVICE_SUCCESS,
			constants.MSG_SUCCESS,
			[]string{},
			nil,
		),
	)
}

func (c *Ctx) Deleted() error {
	return c.JSON(
		http.StatusAccepted,
		c.BuildResponse(
			constants.STATUS_CODE_DELETE_SUCCESS,
			constants.MSG_SUCCESS,
			[]string{},
			nil,
		),
	)
}

func (c *Ctx) BuildResponse(code, message string, errors []string, payload interface{}) *ApiResponse {
	return &ApiResponse{
		Code:    code,
		Message: message,
		Errors:  errors,
		Payload: payload,
	}
}

func (c *Ctx) BindAndValidate(f interface{}) error {
	if err := c.Bind(f); err != nil {
		return err
	}
	if err := c.Validate(f); err != nil {
		return err
	}
	return nil
}

func (c *Ctx) ValidationErrors(errs error) *ApiResponse {
	payload := []FieldValidationError{}
	for _, err := range errs.(validator.ValidationErrors) {
		errObj := &FieldValidationError{}
		errObj.Field = err.Field()
		errObj.Namespace = err.Namespace()
		errObj.Kind = err.Kind().String()
		errObj.Value = err.Value()
		errObj.Error = fmt.Sprintf("%s %s", err.Tag(), err.Param())
		payload = append(payload, *errObj)
	}
	if config.DevModeFlag {
		str, _ := json.MarshalIndent(payload, "", "  ")
		c.Logger().Error("ValidationErrors:")
		c.Logger().Error(string(str))
	}
	return c.BuildResponse(
		constants.STATUS_CODE_VALIDATION_ERROR,
		constants.MSG_VALIDATION_ERROR,
		[]string{constants.MSG_VALIDATION_ERROR},
		payload,
	)
}
