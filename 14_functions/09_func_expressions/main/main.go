package main

import (
	"fmt"

	"github.com/pawarvishal/golangtraining/14_functions/09_func_expressions/greeter"
)

func main() {
	name := "vishal"

	greeting := func(name string) { //anonymous function assigned to identifier, only way to create fn inside another function
		fmt.Println("hello", name)
	}

	greeting(name)

	fmt.Printf("%T \n", greeting)

	greet := greeter.Makegreeter()

	fmt.Println(greet())

	fmt.Printf("%T \n", greet)

}
