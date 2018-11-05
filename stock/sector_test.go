package stock

import (
	"testing"

	"github.com/rendon/testcli"
)

func TestSector(t *testing.T) {
	testcli.Run("../iex-cli", "sector")

	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}

	if !testcli.StdoutContains("NAME") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "NAME")
	}

	if !testcli.StdoutContains("PERFORMANCE") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "PERFORMANCE")
	}

	if !testcli.StdoutContains("LAST UPDATED") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "LAST UPDATED")
	}
}
