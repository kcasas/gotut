package main

// Package level variables
var id int

/*
	Contants
	- only single way to initialize
	- cannot be changed, of course
	- variable name doesn't have to be uppercased
*/
const Pi = 3.14

func main() {
	// Function level variables

	// - constant
	const foreverExists = false

	// Single
	// - variable declaration
	var firstname string
	// - variable initialization
	var lastname string = "Casas"
	// - variable initialization w/ type inference
	language := "tagalog"
	// - assignment
	language = "waray"

	// Multiple
	// - variable declaration
	var online, employed bool
	// - variable initialization
	var age, weight int = 20, 70
	// - variable initialization w/ type inference
	religion, height := "catholic", 12
	// - assignment
	online, employed = true, false

	/*
		Illegal statements

		- statically typed, variables cannot change data type
		online = "false"

		- variable name cannot start with a number
		var 4ever = false

		- of course constants cannot be changed
		foreverExists = true
	*/

	/*
		- basic data types https://tour.golang.org/basics/11
		- zero values https://tour.golang.org/basics/12
	*/

	/*
		- Go requires all variables to be used
		- constants need not be used
		- `_` usually means to ignore in Go
		- this is a workaround to trick Go that we've used the other variables
	*/
	_, _, _, _, _, _, _, _, _ =
		firstname,
		lastname,
		online,
		employed,
		language,
		age,
		weight,
		religion,
		height

}
