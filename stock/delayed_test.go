package stock

import (
	"testing"

	"github.com/rendon/testcli"
)

func TestDelayed(t *testing.T) {
	testcli.Run("../iex-cli", "delayed", "AAPL")

	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}

	if !testcli.StdoutContains("Delayed Price") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Delayed Price")
	}

	if !testcli.StdoutContains("High") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "High")
	}

	if !testcli.StdoutContains("Low") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Low")
	}

	if !testcli.StdoutContains("Delayed Size") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Delayed Size")
	}

	if !testcli.StdoutContains("Delayed Price Time") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Delayed Price Time")
	}
}

func TestDelayedNoArgs(t *testing.T) {
	testcli.Run("../iex-cli", "delayed")

	if !testcli.StdoutContains("Error: No argument supplied") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Error: No argument supplied")
	}
}
