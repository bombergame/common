package errs

import (
	"testing"
)

func TestConstraintViolationError(t *testing.T) {
	if NewConstraintViolationError("constraint violation") == nil {
		t.Error("error not created")
	}
}
