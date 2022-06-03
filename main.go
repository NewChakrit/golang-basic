package main

import "fmt"

// var H = 100  -- Global variable --cannot write type inferrence**

func main() {
	println("Hello World")

	// package "fmt"
	// fmt.Println("Hello scound world")

	fmt.Printf("Hello %d \n", 10)

	fmt.Printf("Hello %v \n", true)     //%v print value
	fmt.Printf("Hello %v \n", "World")  //%v print value
	fmt.Printf("Hello %#v \n", "World") //%v print > value  -- more => Hello "World"

	fmt.Printf("%T \n", false) //%T check type

	// ---- Declaretion ----
	// var -- can change value
	// const -- cannot change value

	var q int
	_ = q // declared but never use

	var x int = 10
	fmt.Println(x)

	//  -- type inferrence --
	// xx := 9
	xx := "New"
	fmt.Println(xx)

	// **** Every declaretion has defuault value (we call zero value!!) ****
	// Zero value
	// string : ""
	// int : 0
	// float : 0
	// bool : false

	//exam
	var xxx int
	var xxxx string //or
	a := 0
	aa := false
	fmt.Println(xxx)  // 0
	fmt.Println(xxxx) // "" (empty)
	fmt.Println(a)    // "" (empty)
	fmt.Println(aa)   // "" (empty)

	fmt.Println("\n")
	fmt.Println("\n")

}
