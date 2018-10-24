package errors

import (
	"testing"
)

func TestNotFoundError(t *testing.T) {
	if NewNotFoundError("not found") == nil {
		t.Error("error not created")
	}
}
