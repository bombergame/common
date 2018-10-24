package log

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger()
	if logger == nil {
		t.Error("logger not created")
	}

	if logger.AsLogrusLogger() == nil {
		t.Error("logger cast failed")
	}
}
