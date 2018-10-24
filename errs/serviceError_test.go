package errs

import (
	"errors"
	"testing"
)

func TestServiceError(t *testing.T) {
	err := NewServiceError(nil)
	if err == nil {
		t.Error("service error not created")
	}

	if err.Error() != ServiceErrorMessage {
		t.Error("wrong error message")
	}

	sErr, ok := err.(*ServiceError)
	if !ok {
		t.Error("cast error -> ServiceError fail")
	}

	if sErr.InnerError() != "" {
		t.Error("wrong error message")
	}
}

func TestServiceInnerError(t *testing.T) {
	const errMsg = "std error"

	err := NewServiceError(errors.New(errMsg))
	if err == nil {
		t.Error("service error not created")
	}

	if err.Error() != ServiceErrorMessage {
		t.Error("wrong error message")
	}

	sErr, ok := err.(*ServiceError)
	if !ok {
		t.Error("cast error -> ServiceError fail")
	}

	if sErr.InnerError() != errMsg {
		t.Error("wrong error message")
	}
}
