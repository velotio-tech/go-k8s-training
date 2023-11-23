package exception

// STATUS_ERROR_TYPE is the type of all premier error messages of exception class
type STATUS_ERROR_TYPE string

const (
	STATUS_INCOMPLETE_PARAM          STATUS_ERROR_TYPE = "incomplete_param"
	STATUS_INVALID_PARAM_TYPE        STATUS_ERROR_TYPE = "invalid_param_type"
	STATUS_INVALID_DATA              STATUS_ERROR_TYPE = "invalid_data"
	STATUS_DB_ERROR                  STATUS_ERROR_TYPE = "db_error"
	STATUS_DB_UNAVAILABLE            STATUS_ERROR_TYPE = "db_unavailable"
	STATUS_UNSUPPORTED_MEDIA_TYPE    STATUS_ERROR_TYPE = "unsupported_media_type"
	STATUS_INVALID_PATH_TYPE         STATUS_ERROR_TYPE = "invalid_path_type"
	STATUS_INCOMPLETE_PATH_TYPE      STATUS_ERROR_TYPE = "incomplete_path_value"
	STATUS_INTERNAL_OPERATION_FAILED STATUS_ERROR_TYPE = "internal_operation_failed"
	STATUS_INVALID_PAYLOAD           STATUS_ERROR_TYPE = "payload_validation_error"
	STATUS_MARSHALLING_ERROR         STATUS_ERROR_TYPE = "marshaling_error"
	STATUS_UNKNOWN_ERROR             STATUS_ERROR_TYPE = "unknown_error"
	STATUS_UNAUTHORIZED_ACCESS       STATUS_ERROR_TYPE = "unauthorized_access"
	STATUS_INVALID_HEADER            STATUS_ERROR_TYPE = "invalid_header"
	STATUS_SERVICE_ERROR             STATUS_ERROR_TYPE = "service_error"
)

// Exception ...
type Exception struct {
	Err        error             `json:"-"`
	Message    string            `json:"message"`
	StatusText STATUS_ERROR_TYPE `json:"status"`
	StatusCode int               `json:"code"`
	Trace      string            `json:"-"`
}
