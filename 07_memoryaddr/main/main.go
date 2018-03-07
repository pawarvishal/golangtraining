package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	ans := number + 10

	fmt.Println("ans is ", ans)
	fmt.Println("address for ans in memory ", &ans)
	fmt.Println("address for number in memory ", &number)

}
