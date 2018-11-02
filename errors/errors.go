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

// ErrorNoArgs handles errors with no stack trace
func ErrorNoArgs() {
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("%s\n", red("Error: No argument supplied"))
}

// ErrorUnmarshal handles errors with no stack trace
func ErrorUnmarshal(err error) {
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("%s\n", red("Error: Failed to unmarshal JSON body"))
	panic(err)
}
