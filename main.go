package main

import (
	"fmt"
	"unicode/utf8"

	"main.go/customer"
	"main.go/user"
)

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

	// ----- If Else -----

	point := 50
	// if point >= 100 {
	// 	fmt.Println(point, "is less than 100")
	// } else if point <= 50 {
	// 	fmt.Println(point, "is more or equl 50")
	// } else {
	// 	fmt.Println("Point is ", point)
	// }

	if point >= 50 && point <= 100 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Not pass")
	}

	fmt.Println("\n")

	// ----- Array -----

	//fix index

	var num [3]int = [3]int{1, 2, 3}
	fmt.Println(num, "\n") //[1, 2, 3]
	fmt.Println(num[0])    // 1
	fmt.Println(num[1])    // 2
	fmt.Println(num[2])    // 3

	// var num2 [3]int = [3]int{} // Zero value
	num2 := [3]int{}          // short declaretion
	fmt.Println(num2)         // [ 0, 0, 0]
	fmt.Printf("%#v\n", num2) //[3]int{0, 0, 0}

	// 2D array
	//syntax first[] = [number of member in array]
	// second[] = [number of index]
	m := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(m[2][2]) //9

	// infinit value array

	num3 := [...]int{1, 2, 3, 4}
	fmt.Println(num3)

	// ---- Slice ----
	// slice can add value in arrray by use func(append)
	num4 := []int{1, 2, 3, 4}
	num4 = append(num4, 5)  //append return num4+ 5
	num5 := append(num4, 6) // num4 + 6 =num5

	checkQuanNum := len(num5)

	fmt.Printf("%#v\n", num4)
	fmt.Println(num4)
	fmt.Println(num5)

	fmt.Println(checkQuanNum) //6
	fmt.Printf("%v, %v\n", num5, checkQuanNum)

	//** Warnning ** Don't use <len> to check quantity of string
	// name1 := "abc"
	// fmt.Println(len(name1)) //3

	// name2 := "ก"
	// fmt.Println(len(name2)) //3 because defference bit

	//if you want to count string, Use uft8.RuneCountInString

	name3 := "ฟหefg"
	fmt.Println(utf8.RuneCountInString(name3)) //5

	//---- set value by length -----

	num6 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	num7 := num6[3:]
	fmt.Println(num7) //40 50 60 70 80 90

	num6 = append(num6, 110)
	num8 := num6[:5]
	fmt.Println(num8)      //10 20 30 40 50
	fmt.Println(num8[2:4]) //30 40

	//  ----- Map -----
	// var countries map[string]string{}

	// countries:= map[string] string{"TH": "Thai", "JP":"Japan"}
	countries := map[string]string{}
	countries["TH"] = "Thailand"
	countries["JP"] = "Japan"

	fmt.Println(countries["TH"]) //Thailand
	fmt.Println(countries)

	// x,y := 10,20 -- mutiple declaretion

	// ---- if else ----
	country, checkCountry := countries["jj"]
	fmt.Println("country is :", country, ",country has value :", checkCountry)

	if checkCountry {
		fmt.Println(country)
	} else {
		fmt.Println("Don't has country")
	}

	// ---- For Loop ----

	value := []int{1, 2, 3, 4, 5, 6}

	for i := 0; i < len(value); i++ {
		fmt.Println(value[i])
	}

	// or

	ii := 0
	for ii < len(value) {
		fmt.Println(value[ii])
		ii++
	}

	fmt.Println("\n")

	// ----- for range -----

	values := []int{90, 80, 70, 60, 50, 40, 30, 20, 10}

	for i, v := range values { // i is index , v is value
		fmt.Println("Index :", i, ", Value :", v)

	}

	//if don't want some value

	for _, v := range values {
		fmt.Println("Value :", v)
	}

	//or
	for i, _ := range values {
		fmt.Println("Index :", i)
	}

	// func sum
	fmt.Println(sum())

	//func sum1
	d, e, f := sum1(4, 5)
	fmt.Println(d, e, f) // 9 hello true

	//if don't want some value
	d, _, _ = sum1(7, 8)
	fmt.Println(d) //15

	fmt.Println("\n")

	//annonymous func --- doesn't has name

	s := func(a, b int) int {
		return a + b
	}

	sum2 := s(10, 20)

	fmt.Println(sum2) //30

	//func advance

	cal(add) // 60
	cal(sub) //40
	// cal(hello) //error  signature type wrong

	g := func(a, b int) int {
		return a * b
	}
	cal(g) //500

	//finally
	cal(func(a, b int) int {
		return a / b //5
	})

	//array parameter
	v := []int{1, 2, 3}
	k := sum3(v)
	fmt.Println(k) // 6

	//validic slice parameter
	k = sum4(5, 6, 7, 8)
	fmt.Println(k) //26

	// ----- Package -----
	fmt.Println(user.Name)     // User
	fmt.Println(customer.Name) //New

	customer.Customer() //Hello Customer

	fmt.Println(customer.Hello()) //Hello Chakrit

}

//  ----- Func -----

func sum() int {
	a := 10
	b := 20
	return a + b

}

func sum1(a, b int) (int, string, bool) {
	return a + b, "Hello", true
}

// ----- annonymous func -----

// type (int,int) int
func add(a, b int) int {
	return a + b
}

// type (int,int) int
func sub(a, b int) int {
	return a - b
}

//type (string) string
func hello(name string) string {
	return "Hello " + name
}

func cal(f func(int, int) int) {
	sum := f(50, 10)
	fmt.Println(sum)
}

// ----- mutiparameter func (array parameter) -----

func sum3(a []int) int {
	s := 0
	for _, v := range a {
		s += v
	}

	return s
}

// ----- func validic slice parameter -----

func sum4(a ...int) int {
	s := 0
	for _, v := range a {
		s += v
	}

	return s
}
