package errors

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/pkg/errors"
)

var red = color.New(color.FgRed).SprintFunc()

// Error handles errors
func Error(description string) {
	err := errors.New(red("Error: " + description))
	fmt.Printf("%+v\n", err)
	os.Exit(1)
}
