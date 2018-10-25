package errs

import (
	"testing"
)

func TestInvalidFormatError(t *testing.T) {
	if NewInvalidFormatError("invalid format") == nil {
		t.Error("error not created")
	}
}
