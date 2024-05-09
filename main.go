package main

import "fmt"

func closure() func() { //returns a function
	text := "Inside the closure function"
	function := func() { // just print text on the screen
		fmt.Println(text)
	}
	return function
}

func main() {
	text := "Inside the main function"
	fmt.Println(text)

	funcN := closure()
	funcN()
}
