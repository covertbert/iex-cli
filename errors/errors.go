package errors

import (
	"fmt"

	"github.com/fatih/color"
)

// Error handles errors
func Error(err error, description string) {
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("\n%s\n\n", red(description))

	panic(err)
}
