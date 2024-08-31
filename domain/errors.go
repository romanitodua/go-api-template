package domain

const (
	ErrTypeValidation  = "validation"  // includes bad request
	ErrTypePermissions = "permissions" // includes forbidden, unauthorized
	ErrTypeNotFound    = "not found"   // includes not found
	ErrTypeTimeout     = "timeout"     // includes timeout
)

// Error messages
const (
	ErrorNotFound       = "nothing was found"
	ErrorForbidden      = "forbidden"
	ErrorUnauthorized   = "unauthorized"
	ErrorBadRequest     = "bad request"
	ErrorValidation     = "validation error"
	ErrorExpired        = "expired"
	ErrorSomethingWrong = "something went wrong"
	ErrorUnknownPath    = "path is not determined"
	ErrorInvalidJson    = "invalid json payload"
)

type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}
func NewError(errorType, message string) *Error {
	return &Error{
		Type:    errorType,
		Message: message,
	}
}
func NewInvalidError(text string) *Error {
	return NewError(ErrTypeValidation, text)
}

func NewForbiddenError(text string) *Error {
	return NewError(ErrTypePermissions, text)
}
func NewInvalidJsonError() *Error {
	return NewError(ErrTypeValidation, ErrorInvalidJson)
}
func NewNotFoundError(text string) *Error {
	return NewError(ErrTypeNotFound, text)
}
func NewTimeoutError(text string) *Error {
	return NewError(ErrTypeTimeout, text)
}
