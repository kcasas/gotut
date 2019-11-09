package main

import "fmt"

func ifelse() {
	/*
		Go's if statement is similar to other languages
	*/
	age := 20
	if age < 12 {
		fmt.Println("boy")
	} else if age > 19 {
		fmt.Println("adult")
	} else {
		fmt.Println("teenager")
	}
}
