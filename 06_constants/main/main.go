package main

import (
	"fmt"

	"github.com/pawarvishal/golangtraining/06_constants/declareconsts"

	"github.com/pawarvishal/golangtraining/06_constants/iotas"
)

const p string = "upskill or perish"

//p := "dont upskill"

func main() {

	const q = "true story"

	fmt.Println("P is ", p)
	fmt.Println("q is ", q)
	fmt.Println(declareconsts.Pi)

	fmt.Printf("KB in Binary %b \n", iotas.KB)
	fmt.Printf("KB in Decimal %d \n", iotas.KB)

	fmt.Printf("GB in Binary %b \n", iotas.GB)
	fmt.Printf("GB in Decimal %d \n", iotas.GB)

}
