package main

import "fmt"

func main() {
	println("Hello World")

	// package "fmt"
	// fmt.Println("Hello scound world")

	fmt.Printf("Hello %d \n", 10)

	fmt.Printf("Hello %v \n", true)     //%v print value
	fmt.Printf("Hello %v \n", "World")  //%v print value
	fmt.Printf("Hello %#v \n", "World") //%v print > value  -- more => Hello "World"

	fmt.Printf("%T \n", false) //%T check type

}
