package logs

import (
	"bou.ke/monkey"
	"github.com/sirupsen/logrus"
	"reflect"
	"testing"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger()
	if logger == nil {
		t.Error("logger not created")
	}

	infoCalled := false
	monkey.PatchInstanceMethod(reflect.TypeOf(logger.logger), "Info",
		func(_ *logrus.Logger, _ ...interface{}) {
			infoCalled = true
		},
	)

	logger.Info("info")
	if !infoCalled {
		t.Error("info method not called")
	}

	errorCalled := false
	monkey.PatchInstanceMethod(reflect.TypeOf(logger.logger), "Error",
		func(_ *logrus.Logger, _ ...interface{}) {
			errorCalled = true
		},
	)

	logger.Error("error")
	if !errorCalled {
		t.Error("error method not called")
	}

	fatalCalled := false
	monkey.PatchInstanceMethod(reflect.TypeOf(logger.logger), "Fatal",
		func(_ *logrus.Logger, _ ...interface{}) {
			fatalCalled = true
		},
	)

	logger.Fatal("fatal")
	if !fatalCalled {
		t.Error("fatal method not called")
	}

	if logger.AsLogrusLogger() == nil {
		t.Error("logger cast failed")
	}
}
