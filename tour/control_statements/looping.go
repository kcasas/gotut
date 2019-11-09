package main

func looping() {
	/*
		Looping
		- Go has only one looping construct, the `for` loop
		- `for` loop has three components separated by semicolons
		  	- init statement: executed before the first iteration
			- condition expression: evaluated before every iteration
			- post statement: executed at the end of every iteration
		- every component is optional
	*/
	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
	}

	// Go's version of while
	for sum < 100 {
		sum += sum
	}

	// Go's version of do while
	for {
		sum += sum
		if sum > 1000 {
			break
		}
	}
}
