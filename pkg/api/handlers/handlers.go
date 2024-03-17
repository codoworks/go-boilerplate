/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/codoworks/go-boilerplate/pkg/clients/logger"
	"github.com/codoworks/go-boilerplate/pkg/config"
	"github.com/codoworks/go-boilerplate/pkg/utils/constants"

	"github.com/go-playground/validator/v10"
)

func Success(payload interface{}) *ApiResponse {
	return BuildResponse(
		constants.STATUS_CODE_SERVICE_SUCCESS,
		constants.MSG_SUCCESS,
		[]string{},
		payload)
}

func Accepted() *ApiResponse {
	return BuildResponse(
		constants.STATUS_CODE_SERVICE_SUCCESS,
		constants.MSG_SUCCESS,
		[]string{},
		nil)
}

func Deleted() *ApiResponse {
	return BuildResponse(
		constants.STATUS_CODE_DELETE_SUCCESS,
		constants.MSG_SUCCESS,
		[]string{},
		nil)
}

func ValidationErrors(errs error) *ApiResponse {
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
		logger.Error("ValidationErrors:")
		logger.Error(string(str))
	}
	return BuildResponse(
		constants.STATUS_CODE_VALIDATION_ERROR,
		constants.MSG_VALIDATION_ERROR,
		[]string{constants.MSG_VALIDATION_ERROR},
		payload)
}
