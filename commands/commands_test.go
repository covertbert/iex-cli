package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	assert.NotZero(t, CliCommands)
}
