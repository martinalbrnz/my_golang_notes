package main

import "fmt"

var myNum int

func main() {
	fmt.Println("--Types--")
	myNum = 36
	aString := "Hello"
	anotherString := "world!"
	const constantInteger int8 = 123
	var aBoolean = true
	aFloat := 3.46
	fmt.Printf("Types of vars: %T %T %T %T %T\n", myNum, aString, aBoolean, anotherString, aFloat)
	// console: Types of vars: int string bool string

	var naturalNullString string
	var naturalNullInt int
	fmt.Println(naturalNullInt, naturalNullString) // 0 ""
	naturalNullInt += 2
	naturalNullString += "Good"
	fmt.Println(naturalNullInt, naturalNullString) // 2 Good

	// Lists
	var catNames = [4]string{"Pol", "Snow", "Daisy", "Orlando"}

	for i := 0; i < len(catNames); i++ {
		fmt.Println("Cat number", i+1, "is", catNames[i])
	}

}
