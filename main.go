package main

import "fmt"

func main() {
	println("Hello World")

	//prinln ปริ้นค่าง่ายๆ
	fmt.Println("Hello Chakrit")

	// %d บอกเลขฐาน 10
	fmt.Printf("Hi %d\n", 10)

	// %v อ่าน value ได้ทุก type
	fmt.Printf("Hi %v\n", true)
	fmt.Printf("Hi %v\n", "Golang")

	// %#v อ่าน > value
	fmt.Printf("Hi %#v\n", "Golang")

	// %T บอก type
	fmt.Printf("Hi %T\n", "Golang")

}
