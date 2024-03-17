/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package errors

import (
	"errors"
	"net/http"

	"github.com/codoworks/go-boilerplate/pkg/api/handlers"
	"github.com/codoworks/go-boilerplate/pkg/utils/constants"

	"github.com/labstack/echo/v4"
)

func AutomatedHttpErrorHandler() echo.HTTPErrorHandler {
	h := &httpErrorHandler{
		statusCodes: ErrorsMap(),
	}
	return h.Handler
}

// func NewHttpErrorHandler(errorStatusCodeMaps map[error]interface{}) *httpErrorHandler {
// 	return &httpErrorHandler{
// 		statusCodes: errorStatusCodeMaps,
// 	}
// }

type httpErrorHandler struct {
	statusCodes map[error]interface{}
}

func (h *httpErrorHandler) getStatusCode(err error) (int, string) {
	for key, value := range h.statusCodes {
		if errors.Is(err, key) {
			val := value.(map[string]interface{})
			return val[constants.WORD_INTERNAL_CODE].(int), val[constants.WORD_SERVICE_CODE].(string)
		}
	}
	return http.StatusInternalServerError, constants.STATUS_CODE_FAILED_TO_DECODE_VALUE
}

func unwrapRecursive(err error) error {
	var originalErr = err
	for originalErr != nil {
		var internalErr = errors.Unwrap(originalErr)
		if internalErr == nil {
			break
		}
		originalErr = internalErr
	}
	return originalErr
}

func (h *httpErrorHandler) Handler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	var serviceCode string
	if ok {
		if he.Internal != nil {
			if herr, ok := he.Internal.(*echo.HTTPError); ok {
				he = herr
			}
		}
	} else {
		var internalCode int
		internalCode, serviceCode = h.getStatusCode(err)
		he = &echo.HTTPError{
			Code:    internalCode,
			Message: unwrapRecursive(err).Error(),
		}
	}

	code := he.Code
	message := he.Message
	if _, ok := he.Message.(string); ok {
		message = handlers.BuildResponse(
			serviceCode,
			constants.MSG_ERROR,
			[]string{err.Error()},
			nil)
	}
	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(he.Code)
		} else {
			err = c.JSON(code, message)
		}
		if err != nil {
			c.Echo().Logger.Error(err)
		}
	}
}
