package main

import (
	"fmt"
)

func main() {
	func(n int) { //anonymous self executing function
		fmt.Println("hello ", n)
	}(10)
}
