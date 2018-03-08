package main

import "fmt"

func main() {
	m := make([]string, 1, 10)
	fmt.Println(m)

	changeMe(m)
	fmt.Println(m[0])
}

func changeMe(s []string) {
	s[0] = "vishal"
	fmt.Println(s)
}
