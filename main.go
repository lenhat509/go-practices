package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println("Hello World!")
	// var strOne string = "string"
	// var strTwo = "string2"
	// var strThree string

	// // short assignment, type inference
	// numberOne := 23

	// fmt.Println(strOne, numberOne, strTwo, strThree)

	// // same line assignment
	// var v1, v2 = true, 123

	// fmt.Println(v1, v2)

	// // type casting
	// v1 := 34.53
	// var v2 int
	// v2 = int(v1)

	// fmt.Println(v2)

	// const v1 = "constant" // can not use short assignment

	// var v1 = 10
	// const v2 = 40  // constants must be computed at compile time, not run-time

	//s1 := true
	// s2 := 2123

	// fmt.Printf("%b", s2)

	// a, b := addAndSubtract(1, 3)

	// fmt.Print(a, b)

	// Array
	// var arrInt [4]int = [4]int{1, 2, 3, 4}
	// arrStr := [3]string{"1", "2", "3"}
	// fmt.Print(arrInt, len(arrInt))

	// // Slices
	// var arrInt = []int{1, 2, 3, 4}
	// fmt.Println(arrInt, len(arrInt))
	// // Add items into a slice
	// arrInt = append(arrInt, 5, 6)
	// fmt.Println(arrInt, len(arrInt))
	// // Slice Range
	// arrIntRange := arrInt[0:3]
	// fmt.Printf("%T", arrIntRange) // print Type of object, can be accomplished with reflect pkg

	result := strings.ReplaceAll("Hello World Hello", "Hello", "World")

	fmt.Println(result)

}

// func addAndSubtract(a, b int) (total, diff int) {
// 	total = a + b
// 	diff = a - b
// 	return 0, 1
// }
