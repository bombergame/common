package errs

type ConstraintViolationError struct {
	ServiceError
}

func NewConstraintViolationError(message string) error {
	return NewConstraintViolationInnerError(message, nil)
}

func NewConstraintViolationInnerError(message string, innerErr error) error {
	err := &ConstraintViolationError{}
	err.message = message
	err.innerErr = innerErr
	return err
}
