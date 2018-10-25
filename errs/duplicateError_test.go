package errs

import (
	"testing"
)

func TestDuplicateError(t *testing.T) {
	if NewDuplicateError("already exists") == nil {
		t.Error("error not created")
	}
}
