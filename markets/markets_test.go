package markets

import (
	"testing"

	"github.com/rendon/testcli"
)

func TestMarkets(t *testing.T) {
	testcli.Run("../iex-cli", "markets")

	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}

	if !testcli.StdoutContains("MIC") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "MIC")
	}

	if !testcli.StdoutContains("VOLUME") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "VOLUME")
	}

	if !testcli.StdoutContains("TAPE A") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "TAPE A")
	}

	if !testcli.StdoutContains("MARKET PERCENT") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "MARKET PERCENT")
	}
}
