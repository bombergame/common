package args

import (
	"os"
	"regexp"
	"testing"
)

func TestGetString(t *testing.T) {
	os.Args = []string{
		"--flag",
		"--address=localhost:8080",
	}

	if GetString("address", "localhost") != "localhost:8080" {
		t.Error("argument not parsed")
	}

	if GetString("flag", "default") != "default" {
		t.Error("default value not returned")
	}

	if GetString("someVar", "default") != "default" {
		t.Error("default value not returned")
	}
}

func TestGetFlag(t *testing.T) {
	os.Args = []string{
		"--flag",
		"--address=localhost:8080",
	}

	if !GetFlag("flag", false) {
		t.Error("flag not parsed")
	}

	if GetFlag("address", false) {
		t.Error("wrong flag parsed")
	}
}

func TestParseGroups(t *testing.T) {
	r, _ := regexp.Compile(`^(?P<name>\w+)$`)

	str := "user"

	mp := parseGroups(r, str)
	if len(mp) != 2 {
		t.Error("wrong number of groups")
	}

	if mp[""] != str || mp["name"] != str {
		t.Error("wrong group value")
	}

	str = "--user"

	mp = parseGroups(r, str)
	if mp != nil {
		t.Error("wrong number of groups")
	}
}
