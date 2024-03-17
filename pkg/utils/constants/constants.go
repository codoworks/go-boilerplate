/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package constants

import (
	"errors"
)

const (
	// defaults
	DEFAULT_PROTECTED_API_PORT       string = "8080"
	DEFAULT_PUBLIC_API_PORT          string = "8081"
	DEFAULT_HIDDEN_API_PORT          string = "8079"
	DEFAULT_HOST                     string = "0.0.0.0"
	DEFAULT_DEV_HOST                 string = "127.0.0.1"
	DEFAULT_LOG_LEVEL                string = "warn"
	DEFAULT_DEV_LOG_LEVEL            string = "debug"
	DEFAULT_CORS_ALLOW_ORIGINS       string = "*"
	DEFAULT_DB_PLATFORM              string = DB_PLATFORM_SQLITE
	DEFAULT_DB_TIMEZONE              string = DB_TIMEZONE_UTC
	DEFAULT_DB_SSL_MODE              string = DB_SSL_MODE_DISABLED
	DEFAULT_SQLITE_DB_NAME           string = "sqlite.db"
	DEFAULT_LOGGER_TIMESTAMP_FORMAT  string = "2006-01-02 15:04:05.00000"
	DEFAULT_REQUEST_TIMEOUT_DURATION int    = 60   // seconds
	DEFAULT_WATCHER_SLEEP_INTERVAL   int    = 5000 // milliseconds
	DEFAULT_GZIP_LEVEL               int    = 5

	// features
	FEATURE_SERVICE    string = "service"
	FEATURE_ORY_KRATOS string = "ory_kratos"
	FEATURE_ORY_KETO   string = "ory_keto"
	FEATURE_DATABASE   string = "database"
	FEATURE_CORS       string = "cors"
	FEATURE_GZIP       string = "gzip"
	FEATURE_REDIS      string = "redis"

	// generic words
	WORD_DATABASE        string = "database"
	WORD_DATABASE_SERVER string = "database_server"
	WORD_INTERNAL_CODE   string = "internalCode"
	WORD_SERVICE_CODE    string = "serviceCode"

	// names, specific to this application
	NAME_HEALTH_PATH       string = "/alive"
	NAME_HEALTH_READY_PATH string = "/ready"
	NAME_CORS_CONFIG       string = "CORSAllowOrigins"
	NAME_TIMEOUT_DURATION  string = "RequestTimeoutDuration"
	NAME_CMD_DB_MIGRATE    string = "migrate"
	NAME_CMD_DB_ROLLBACK   string = "rollback"
	NAME_CMD_DB_SEED       string = "seed"

	DB_PLATFORM_POSTGRES  string = "postgres"
	DB_PLATFORM_MYSQL     string = "mysql"
	DB_PLATFORM_SQLITE    string = "sqlite"
	DB_SSL_MODE_ENABLED   string = "require"
	DB_SSL_MODE_DISABLED  string = "disable"
	DB_TIMEZONE_UTC       string = "Etc/GMT"
	DB_TIMEZONE_MELBOURNE string = "Australia/Melbourne"

	HEADER_CONTENT_TYPE      string = "Content-Type"
	HEADER_CONTENT_TYPE_JSON string = "application/json; charset=UTF-8"
	HEADER_AUTHORIZATION     string = "Authorization"
	HEADER_AUTH_BEARER_WORD  string = "Bearer"
	HEADER_KRATOS_COOKIE     string = "ory_kratos_session"

	// output messages
	MSG_SERVER_SHUTTING_DOWN        string = "server is shutting down"
	MSG_NOT_ACCEPTABLE              string = "not acceptable"
	MSG_MISSING_ACCEPT_HEADER       string = "unknown accept format"
	MSG_MISSING_CONTENT_TYPE_HEADER string = "unknown content format"
	MSG_SUCCESS                     string = "success"
	MSG_ERROR                       string = "error"
	MSG_VALIDATION_ERROR            string = "validation error"
	MSG_ROUTE_NOT_FOUND             string = "route not found"
	MSG_RECORD_NOT_FOUND            string = "record not found"
	MSG_DEPENDENCY_NOT_FOUND        string = "dependency not found"
	MSG_SESSION_NOT_FOUND           string = "session not found"
	MSG_ACCESS_IDS_NOT_FOUND        string = "access ids not found or not readable"
	MSG_NOT_AUTHORIZED              string = "not authorized"
	MSG_ID_NOT_READABLE             string = "ID not found or not readable"
	MSG_UNABLE_TO_BIND_BODY         string = "error binding body"
	MSG_FORBIDDEN                   string = "forbidden"
	MSG_UNKNOWN_DB_PLATFORM         string = "unknown database platform"

	// output status codes
	STATUS_CODE_SERVICE_SUCCESS                            string = "20001"
	STATUS_CODE_DELETE_SUCCESS                             string = "20201"
	STATUS_CODE_ERROR_BINDING_BODY                         string = "40002"
	STATUS_CODE_INVALID_TRANSACTION                        string = "40003"
	STATUS_CODE_PRIMARY_KEY_REQUIRED                       string = "40004"
	STATUS_CODE_MODEL_VALUE_REQUIRED                       string = "40005"
	STATUS_CODE_UNSUPPORTED_DRIVER                         string = "40006"
	STATUS_CODE_REGISTERED                                 string = "40007"
	STATUS_CODE_INVALID_FIELD                              string = "40008"
	STATUS_CODE_INVALID_DATA                               string = "40009"
	STATUS_CODE_INVALID_DB                                 string = "40010"
	STATUS_CODE_INVALID_VALUE                              string = "40011"
	STATUS_CODE_NOT_IMPLEMENTED                            string = "40012"
	STATUS_CODE_MISSING_WHERE_CLAUSE                       string = "40013"
	STATUS_CODE_UNSUPPORTED_RELATION                       string = "40014"
	STATUS_CODE_EMPTY_SLICE                                string = "40015"
	STATUS_CODE_DRY_RUN_UNSUPPORTED                        string = "40016"
	STATUS_CODE_INVALID_VALUE_LENGTH                       string = "40017"
	STATUS_CODE_PRELOAD_NOT_ALLOWED                        string = "40018"
	STATUS_CODE_VALIDATION_ERROR                           string = "40019"
	STATUS_CODE_IDS_NOT_READABLE                           string = "40020"
	STATUS_CODE_NOT_AUTHORIZED_WITHOUT_HEADER              string = "40101"
	STATUS_CODE_NOT_AUTHORIZED                             string = "40102"
	STATUS_CODE_ROUTE_NOT_FOUND                            string = "40401"
	STATUS_CODE_RECORD_NOT_FOUND                           string = "40402"
	STATUS_CODE_DEPENDENCY_NOT_FOUND                       string = "40403"
	STATUS_CODE_ID_NOT_FOUND                               string = "40404"
	STATUS_CODE_NOT_ACCEPTABLE_WITHOUT_ACCEPT_HEADER       string = "40601"
	STATUS_CODE_NOT_ACCEPTABLE_WITHOUT_CONTENT_TYPE_HEADER string = "40602"
	STATUS_CODE_FAILED_TO_DECODE_VALUE                     string = "50011"
)

var (
	// log levels
	LOG_LEVELS []string = []string{"debug", "info", "warn", "error", "fatal", "panic"}

	// custom errors
	ERROR_NOT_AUTHORIZED       = errors.New(MSG_NOT_AUTHORIZED)
	ERROR_SESSION_NOT_FOUND    = errors.New(MSG_SESSION_NOT_FOUND)
	ERROR_ID_NOT_FOUND         = errors.New(MSG_ACCESS_IDS_NOT_FOUND)
	ERROR_ACCESS_IDS_NOT_FOUND = errors.New(MSG_SESSION_NOT_FOUND)
	ERROR_BINDING_BODY         = errors.New(MSG_UNABLE_TO_BIND_BODY)
	ERROR_UNKNOWN_DB_PLATFORM  = errors.New(MSG_UNKNOWN_DB_PLATFORM)
)
