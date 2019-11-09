package main

import (
	"fmt"
	"runtime"
)

func switches() {
	/*
		Switch in Go
		- `switch` has two components:
		  - init statement
		  - value to compare the cases with
		- components are optional
		- only runs the first case that evaluates to true. no `break` statement is needed at the end of the case.
		- cases

	*/

	/*
		- syntax is `switch {init statement}; {value}`
		- both {init statement} and {value} are optional
	*/
	linux := "linux"
	switch os := runtime.GOOS; os {
	// cases can be constant
	case "darwin":
		fmt.Println("OS X.")
	// cases can be variable
	case linux:
		fmt.Println("Linux.")
	// cases can be expression
	case "darwin" + linux:
		fmt.Println("darwin" + linux)
	case true:
		fmt.Println("gago")
	default:
		fmt.Printf("%s.\n", os)
	}

	switch {
	case true:
		fmt.Println("darwin")
	}
}
