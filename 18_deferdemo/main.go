package main

import "fmt"

func f1() {
	fmt.Println("In function f1()")
}

func f2() {
	fmt.Println("In function f2()")
}

func f3() {
	fmt.Println("In function f3()")
}

func f4() {
	fmt.Println("In function f4()")
}

func main() {
	defer f3() //defers the excecution of function before main exists ..first defer last executed
	f1()
	defer f2()
	f4()
}
