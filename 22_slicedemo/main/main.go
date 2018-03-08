package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(mySlice)
	fmt.Printf("%T \n", mySlice)

	fmt.Println(mySlice[2:4])
	fmt.Println("Vishal"[0])

	newSlice := make([]int, 0, 5)

	for i := 1; i <= 5; i++ {
		newSlice = append(newSlice, i)
	}

	fmt.Println("Initial slice")
	fmt.Println("length of the slice ", len(newSlice))
	fmt.Println("Capacity of the slice ", len(newSlice))
	fmt.Println(newSlice)

	fmt.Println("After adding extra element beyond capacity")
	newSlice = append(newSlice, 6)
	fmt.Println("length of the slice ", len(newSlice))
	fmt.Println("Capacity of the slice ", cap(newSlice))
	fmt.Println(newSlice)

	fmt.Printf("Type of the slice %T \n", newSlice)

	//appending slice to slice
	mySlice = append(mySlice, newSlice...)
	fmt.Println("after adding both the slices")
	fmt.Println(mySlice)

}
