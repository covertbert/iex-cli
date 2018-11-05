package stock

import (
	"testing"

	"github.com/rendon/testcli"
)

func TestIPO(t *testing.T) {
	testcli.Run("../iex-cli", "ipo")

	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}

	if !testcli.StdoutContains("Company") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Company")
	}

	if !testcli.StdoutContains("Symbol") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Symbol")
	}

	if !testcli.StdoutContains("Price") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Price")
	}

	if !testcli.StdoutContains("Shares") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Shares")
	}

	if !testcli.StdoutContains("Amount") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Amount")
	}
}
