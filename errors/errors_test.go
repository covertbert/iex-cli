package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	assert.Panics(t, func() { Error(errors.New("Chicken"), "Hello") }, "The code did not panic")
}
