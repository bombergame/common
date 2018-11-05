package errs

import (
	"testing"
)

func TestNotAuthorizedError(t *testing.T) {
	if NewNotAuthorizedError() == nil {
		t.Error("error not created")
	}
}
