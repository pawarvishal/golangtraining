package main

import "fmt"

var a = "some string" //package level scope

func main() {

	fmt.Println(a)
	b := "another string" //block level scope

	fmt.Println(b)

	y := 17

	foo(y)
}

func foo(z int) {

	fmt.Println("passed parameter ", z)
	fmt.Println("package level variable ", a)
}
