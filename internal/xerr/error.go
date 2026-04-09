package xerr

type APIError struct {
	Code    int
	Message string
}

func (e *APIError) Error() string {
	return e.Message
}

func NewException(code int, msg string) *APIError {
	return &APIError{
		Code:    code,
		Message: msg,
	}
}
