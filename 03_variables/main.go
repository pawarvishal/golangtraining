package main

import "fmt"

var fname = "vishal"
var lname = "pawar"

func main() {
	a := 10
	b := "some string"
	c := 45.3
	d := true

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)

	fmt.Println("\n", fname)
	fmt.Println("\n", lname)

}
