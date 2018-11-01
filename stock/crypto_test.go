package stock

import (
	"testing"

	"github.com/rendon/testcli"
)

func TestCrypto(t *testing.T) {
	testcli.Run("../iex-cli", "crypto")

	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}

	if !testcli.StdoutContains("Symbol") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Symbol")
	}

	if !testcli.StdoutContains("High") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "High")
	}

	if !testcli.StdoutContains("Low") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Low")
	}

	if !testcli.StdoutContains("Latest") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Latest")
	}
}
