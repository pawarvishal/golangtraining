package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello", name)
}

func main() {
	greet("vishal")
	greet("noone")
}
