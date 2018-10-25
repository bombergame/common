package errs

type InvalidFormatError struct {
	ServiceError
}

func NewInvalidFormatError(message string) error {
	return NewInvalidFormatInnerError(message, nil)
}

func NewInvalidFormatInnerError(message string, innerErr error) error {
	err := &InvalidFormatError{}
	err.message = message
	err.innerErr = innerErr
	return err
}
