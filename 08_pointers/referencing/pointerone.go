package main

import "fmt"

func main() {
	a := 23

	var b = &a

	fmt.Println("value of a ", a)
	fmt.Println("address of a ", &a)

	fmt.Println("value of b ", b)
	fmt.Println("value in b deferencing ", *b)

	*b = 43

	fmt.Println("value in a after modify", a)
}
