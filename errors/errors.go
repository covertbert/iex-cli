package errors

import (
	"fmt"

	"github.com/fatih/color"
)

// Error handles errors
func Error(err error, description string) {
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("\nError: %s\n\n", red(description))

	panic(err)
}

// ErrorNS handles errors with no stack trace
func ErrorNS(description string) {
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("%s\n", red("Error: "+description))
}
