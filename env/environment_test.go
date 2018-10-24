package env

import (
	"os"
	"testing"
)

func genVarName(base string) string {
	for {
		if os.Getenv(base) != "" {
			base += "1"
		} else {
			break
		}
	}
	return base
}

func TestGetSetVar(t *testing.T) {
	envVar1 := genVarName("SOME_ENV_VAR")
	envVar2 := genVarName("ANOTHER_ENV_VAR")

	const envVarValue = "some_value"
	const envVarDefault = "default_value"

	SetVar(envVar1, envVarValue)

	if GetVar(envVar1, envVarDefault) != envVarValue {
		t.Error("environment variable not set")
	}

	if GetVar(envVar2, envVarDefault) != envVarDefault {
		t.Error("default variable value not returned")
	}
}
