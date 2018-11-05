package errors

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/pkg/errors"
)

var red = color.New(color.FgRed).Add(color.Bold).SprintFunc()
var yellow = color.New(color.FgHiYellow).SprintFunc()

// Error handles errors
func Error(description string, errs ...error) {
	err := errors.New(red("\nError: " + description + "\n"))
	fmt.Printf("%+v\n", err)

	for _, e := range errs {
		fmt.Printf("\n%+v\n", yellow(e))
	}

	os.Exit(1)
}
