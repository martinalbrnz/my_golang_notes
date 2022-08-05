package main

import "fmt"

var myNum int

func pointerValueChanger(val *int) {
	fmt.Println("Before change", *val)
	*val = 1
	fmt.Println("Afterchange", *val)
}

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
	// var catNames = [4]string{"Pol", "Snow", "Daisy", "Orlando"}

	// for i := 0; i < len(catNames); i++ {
	// 	fmt.Println("Cat number", i+1, "is", catNames[i])
	// }

	// var someCats = catNames[1:3]

	// for i := 0; i < len(someCats); i++ {
	// 	fmt.Println("Selected cat is", someCats[i])
	// }

	// var allCats = catNames[:]
	// allCats = append(allCats, "Sparkle")

	// for i := 0; i < len(allCats); i++ {
	// 	fmt.Println("A cat is", allCats[i])
	// }

	// var intSlice = []int16{1, 2, 3, 4, 5, 6, 7}

	catsWeight := map[string]float64{
		"Pol": 4.2, "Snow": 5.1, "Daisy": 3.8, "Orlando": 5.2,
	}

	for k, v := range catsWeight {
		fmt.Println(k, "weight is about", v, "k")
	}

	// Reversing keys and values for fun
	keysAndVals := map[string]string{
		"hello": "world", "question": "answer", "day": "night", "cat": "dog",
	}
	var valsAndKeys = make(map[string]string)

	for k, v := range keysAndVals {
		valsAndKeys[v] = k
	}

	for k, v := range valsAndKeys {
		fmt.Println(k, "is not", v)
	}

	aLonelyTwo := 2
	pointerValueChanger(&aLonelyTwo)
}
