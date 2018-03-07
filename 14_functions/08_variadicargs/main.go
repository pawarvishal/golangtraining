package main

import "fmt"

func main() {
	data := []float64{10, 20, 30, 40} //making slice(list)
	n := average(data...)             //sending slice to varidiac function
	fmt.Println("average is ", n)
}

func average(sf ...float64) float64 { //varidiac function
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)

	var total float64

	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
