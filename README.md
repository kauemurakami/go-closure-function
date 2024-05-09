[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-closure-function/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-closure-function/blob/main/README.md)  
go version 1.22.1  

## Função de fechamento
These are functions that refer to variables that are outside their body/scope:  
```go
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
```
Now let's think about what will happen here, in ```closure()``` we are returning a function without a name and without parameters that prints text, but when it is called in ```main()```, it will print the function text or the text in main:  
```go
// outputs
  fmt.Println(text) // Inside the main function
  funcN := closure()
  funcN() // Inside the closure function
```
This is a different behavior from the ```closure()``` function, because when we created the function we were referencing the ```text``` scope variable, so when we execute it, this reference will not be lost, it maintains the initial reference.  
That's what the ```closure``` function is, it has a kind of "memory", where it came from, and we will always be referring to the references passed to it in the function.  