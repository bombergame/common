package errs

import (
	"testing"
)

func TestAccessDeniedError(t *testing.T) {
	if NewAccessDeniedError() == nil {
		t.Error("error not created")
	}
}
