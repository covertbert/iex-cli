package stock

import (
	"testing"

	"github.com/rendon/testcli"
)

func TestNews(t *testing.T) {
	testcli.Run("../iex-cli", "news", "AAPL")

	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}
}

func TestNewsNoArgs(t *testing.T) {
	testcli.Run("../iex-cli", "news")

	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}
}
