package stock

import (
	"testing"

	"github.com/rendon/testcli"
)

func TestPeers(t *testing.T) {
	testcli.Run("../iex-cli", "peers", "AAPL")

	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}
}

func TestPeersNoArgs(t *testing.T) {
	testcli.Run("../iex-cli", "peers")

	if !testcli.StdoutContains("Error: No argument supplied") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Error: No argument supplied")
	}
}
