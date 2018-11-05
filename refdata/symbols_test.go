package refdata

import (
	"testing"

	"github.com/rendon/testcli"
)

func TestSymbols(t *testing.T) {
	testcli.Run("../iex-cli", "symbols")

	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}
}
