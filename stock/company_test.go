package stock

import (
	"testing"

	"github.com/rendon/testcli"
)

func TestCompany(t *testing.T) {
	testcli.Run("../iex-cli", "company", "AAPL")

	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}

	if !testcli.StdoutContains("COMPANY NAME") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "COMPANY NAME")
	}

	if !testcli.StdoutContains("EXCHANGE") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "EXCHANGE")
	}
}
