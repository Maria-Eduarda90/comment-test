package httperr

import "net/http"

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error,omitempty"`
	Code    int      `json:"code"`
	Fields  []Fields `json:"fields,omitempty"`
}

type Fields struct {
	Field   string      `json:"field"`
	Value   interface{} `json:"value,omitempty"`
	Message string      `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(message, err string, code int, field []Fields) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Fields:  field,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, fields []Fields) *RestErr {
	return &RestErr{
		Message: message,
		Err: "bad_request",
		Code: http.StatusBadRequest,
		Fields: fields,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "internal_server_error",
		Code: http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "forbidden",
		Code: http.StatusForbidden,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "not_found",
		Code: http.StatusNotFound,
	}
}

func NewUnauthorizedRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "invalid_credentials",
		Code: http.StatusUnauthorized,
	}
}