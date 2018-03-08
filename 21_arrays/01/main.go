package main

import "fmt"

func main() {
	var x [58]string
	fmt.Println(x)
	fmt.Println(len(x))

	for i := 65; i <= 122; i++ {
		x[i-65] = string(i) //gives back ascii
	}

	fmt.Println(x)
	fmt.Println(x[15])
}
