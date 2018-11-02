package errors

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

// Error handles errors
func Error(description string) {
	hello := errors.New(description)
	fmt.Printf("Error: %+v", hello)
	os.Exit(1)
}
