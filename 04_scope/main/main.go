package main

import (
	"fmt"

	"github.com/pawarvishal/golangtraining/04_scope/vis"
)

var a = "some string" //package level scope

func main() {

	fmt.Println(a)
	b := "another string" //block level scope

	fmt.Println(b)

	y := 17

	foo(y)
	fmt.Println(vis.MyFunc())
}

func foo(z int) {
	fmt.Println(vis.MyName)
	fmt.Println("passed parameter ", z)
	fmt.Println("package level variable ", a)
}
