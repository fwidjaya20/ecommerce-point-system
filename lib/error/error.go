package error

type Error struct {
	Err         error
	StatusCode  int
	MessageCode string
	Message     string
}

func NewError(err error, statusCode int, messageCode string, message string) *Error {
	e := &Error{
		Err:         err,
		StatusCode:  statusCode,
		MessageCode: messageCode,
		Message:     message,
	}

	return e
}

func (e *Error) Error() string {
	return e.Err.Error()
}