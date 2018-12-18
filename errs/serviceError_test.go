package errs

import (
	"errors"
	"testing"

	"github.com/bombergame/common/consts"
)

func TestServiceError(t *testing.T) {
	err := NewInternalServiceError(nil)
	if err == nil {
		t.Error("service error not created")
	}

	if err.Error() != InternalServiceErrorMessage {
		t.Error("wrong error message")
	}

	sErr, ok := err.(*ServiceError)
	if !ok {
		t.Error("cast error -> ServiceError fail")
	}

	if sErr.InnerError() != consts.EmptyString {
		t.Error("wrong error message")
	}

	if sErr.ErrorType() != Internal {
		t.Error("wrong error type")
	}

	const message = "message"
	err = NewInternalServiceError(errors.New("message"))
	sErr, ok = err.(*ServiceError)
	if !ok {
		t.Error("cast error -> ServiceError fail")
	}

	if sErr.InnerError() != message {
		t.Error("wrong error message")
	}
}

func TestConstructors(t *testing.T) {
	const message = "message"
	_ = NewInternalServiceError(nil)
	_ = NewNotFoundError(message)
	_ = NewInvalidFormatError(message)
	_ = NewDuplicateError(message)
	_ = NewBadRequestError(message)
	_ = NewNotAuthorizedError()
	_ = NewAccessDeniedError()
}
