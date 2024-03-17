/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package errors

import (
	"net/http"

	"github.com/codoworks/go-boilerplate/pkg/utils/constants"

	"gorm.io/gorm"
)

func ErrorsMap() map[error]interface{} {
	var errorMap = make(map[error]interface{})

	// service errors
	errorMap[constants.ERROR_ID_NOT_FOUND] = map[string]interface{}{
		"internalCode": http.StatusBadRequest,
		"serviceCode":  constants.STATUS_CODE_ID_NOT_FOUND,
	}
	errorMap[constants.ERROR_ACCESS_IDS_NOT_FOUND] = map[string]interface{}{
		"internalCode": http.StatusBadRequest,
		"serviceCode":  constants.STATUS_CODE_IDS_NOT_READABLE,
	}
	errorMap[constants.ERROR_BINDING_BODY] = map[string]interface{}{
		"internalCode": http.StatusBadRequest,
		"serviceCode":  constants.STATUS_CODE_ERROR_BINDING_BODY,
	}
	errorMap[constants.ERROR_SESSION_NOT_FOUND] = map[string]interface{}{
		"internalCode": http.StatusForbidden,
		"serviceCode":  constants.STATUS_CODE_NOT_AUTHORIZED_WITHOUT_HEADER,
	}
	errorMap[constants.ERROR_NOT_AUTHORIZED] = map[string]interface{}{
		"internalCode": http.StatusForbidden,
		"serviceCode":  constants.STATUS_CODE_NOT_AUTHORIZED,
	}

	// gorm errors
	errorMap[gorm.ErrRecordNotFound] = map[string]interface{}{
		"internalCode": http.StatusNotFound,
		"serviceCode":  constants.STATUS_CODE_RECORD_NOT_FOUND,
	}
	errorMap[gorm.ErrInvalidTransaction] = map[string]interface{}{
		"internalCode": http.StatusBadRequest,
		"serviceCode":  constants.STATUS_CODE_INVALID_TRANSACTION,
	}
	errorMap[gorm.ErrPrimaryKeyRequired] = map[string]interface{}{
		"internalCode": http.StatusBadRequest,
		"serviceCode":  constants.STATUS_CODE_PRIMARY_KEY_REQUIRED,
	}
	errorMap[gorm.ErrModelValueRequired] = map[string]interface{}{
		"internalCode": http.StatusBadRequest,
		"serviceCode":  constants.STATUS_CODE_MODEL_VALUE_REQUIRED,
	}
	errorMap[gorm.ErrUnsupportedDriver] = map[string]interface{}{
		"internalCode": http.StatusBadRequest,
		"serviceCode":  constants.STATUS_CODE_UNSUPPORTED_DRIVER,
	}
	errorMap[gorm.ErrRegistered] = map[string]interface{}{
		"internalCode": http.StatusBadRequest,
		"serviceCode":  constants.STATUS_CODE_REGISTERED,
	}
	errorMap[gorm.ErrInvalidField] = map[string]interface{}{
		"internalCode": http.StatusBadRequest,
		"serviceCode":  constants.STATUS_CODE_INVALID_FIELD,
	}
	errorMap[gorm.ErrInvalidData] = map[string]interface{}{
		"internalCode": http.StatusBadRequest,
		"serviceCode":  constants.STATUS_CODE_INVALID_DATA,
	}
	errorMap[gorm.ErrInvalidDB] = map[string]interface{}{
		"internalCode": http.StatusBadRequest,
		"serviceCode":  constants.STATUS_CODE_INVALID_DB,
	}
	errorMap[gorm.ErrInvalidValue] = map[string]interface{}{
		"internalCode": http.StatusBadRequest,
		"serviceCode":  constants.STATUS_CODE_INVALID_VALUE,
	}
	errorMap[gorm.ErrNotImplemented] = map[string]interface{}{
		"internalCode": http.StatusBadRequest,
		"serviceCode":  constants.STATUS_CODE_NOT_IMPLEMENTED,
	}
	errorMap[gorm.ErrMissingWhereClause] = map[string]interface{}{
		"internalCode": http.StatusBadRequest,
		"serviceCode":  constants.STATUS_CODE_MISSING_WHERE_CLAUSE,
	}
	errorMap[gorm.ErrUnsupportedRelation] = map[string]interface{}{
		"internalCode": http.StatusBadRequest,
		"serviceCode":  constants.STATUS_CODE_UNSUPPORTED_RELATION,
	}
	errorMap[gorm.ErrEmptySlice] = map[string]interface{}{
		"internalCode": http.StatusBadRequest,
		"serviceCode":  constants.STATUS_CODE_EMPTY_SLICE,
	}
	errorMap[gorm.ErrDryRunModeUnsupported] = map[string]interface{}{
		"internalCode": http.StatusBadRequest,
		"serviceCode":  constants.STATUS_CODE_DRY_RUN_UNSUPPORTED,
	}
	errorMap[gorm.ErrInvalidValueOfLength] = map[string]interface{}{
		"internalCode": http.StatusBadRequest,
		"serviceCode":  constants.STATUS_CODE_INVALID_VALUE_LENGTH,
	}
	errorMap[gorm.ErrPreloadNotAllowed] = map[string]interface{}{
		"internalCode": http.StatusBadRequest,
		"serviceCode":  constants.STATUS_CODE_PRELOAD_NOT_ALLOWED,
	}

	return errorMap
}
