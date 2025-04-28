package exception

type NotFoundError struct {
	Error error
}

func NewNotFoundError(err error) NotFoundError {
	return NotFoundError{Error: err}
}