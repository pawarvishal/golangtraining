package main

import "fmt"

func main() {
	fmt.Println("Enter a number to check even or odd ")
	var no int
	fmt.Scan(&no)

	if no % 2 == 0 {
		fmt.Println("no is even ", no)
	} else {
		fmt.Println("no is odd ", no)
	}
}
