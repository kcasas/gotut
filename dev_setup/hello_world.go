/*

GoLang program execution entry point.

- package name should be main
- cannot be imported by other packages
*/
package main

// import core golang package `fmt`
import "fmt"

/*
- package main must declare function main
- takes no arguments and returns no value
- variables and functions can be added outside of the function main
*/
func main() {
	fmt.Println("Hello World!!!")
}

/*
Run: `$ go run main.go`
*/
