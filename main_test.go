package main

import (
	"testing"

	"github.com/rendon/testcli"
)

func TestMain(t *testing.T) {
	testcli.Run("./iex-cli")

	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}

	if !testcli.StdoutContains("NAME") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "NAME")
	}

	if !testcli.StdoutContains("USAGE") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "USAGE")
	}

	if !testcli.StdoutContains("VERSION") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "VERSION")
	}

	if !testcli.StdoutContains("COMMANDS") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "COMMANDS")
	}

	if !testcli.StdoutContains("GLOBAL OPTIONS") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "GLOBAL OPTIONS")
	}
}
