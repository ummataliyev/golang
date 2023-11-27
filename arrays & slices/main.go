package main

import (
	"fmt"
	"sort"
)

func main() {
	testSlices3()
}

func testArrays1() {
	// We create a new array called myarr that contains 3 string elements
	var myarr [3]string

	// When assigning a value to an element, its index is written and indexing starts from 0
	myarr[0] = "GO"
	myarr[1] = "Python"
	myarr[2] = "Java"

	// Index is also used to get the value of the element
	fmt.Println("Elements of the array")
	fmt.Println("1: ", myarr[0])
	fmt.Println("2: ", myarr[1])
	fmt.Println("3: ", myarr[2])
}

// A short way to declare an array
func testArrays2() {

	myarr := [3]int{2, 4, 8}
	fmt.Println(myarr)
}

// Array comparison
func testArrays3() {
	myarr1 := [...]int{9, 7, 6, 5}
	myarr2 := [4]int{9, 7, 6, 5}
	fmt.Println(myarr1 == myarr2)
}

func testArrays4() {
	myarr := [3][3]string{{"C#", "GO", "TypeScript"},
		{"Java", "Python", "PHP"},
		{"C++", "C", "Rust"}}
	fmt.Println(myarr)
}

// Copying arrays
func testArrays5() {
	myarr1 := [3]int{1, 2, 3}
	myarr2 := myarr1
	fmt.Println(myarr1)
	fmt.Println(myarr2)

	myarr1[2] = 100
	fmt.Println(myarr1)
	fmt.Println(myarr2)
}

// Coping arrays with references
func testArrays6() {
	myarr1 := [3]int{2, 4, 8}
	myarr2 := &myarr1
	fmt.Println(myarr1)
	fmt.Println(*myarr2)

	myarr1[2] = 100
	fmt.Println(myarr1)
	fmt.Println(*myarr2)
}

// Slices
func testSlices1() {
	// Declaring new slices
	myslice := []int{2, 4, 8}

	// Append new element to the end of the slice
	myslice = append(myslice, 10)

	// To get the length of slices
	fmt.Printf("Length of slice: %d\n", len(myslice))
}

func testSlices2() {
	myarr := [7]string{"apple", "orange", "granada", "banana",
		"avacado", "ananas", "pear"}

	myslice := myarr[1:4]
	fmt.Println(myslice)
}

func testSlices3() {
	// Sorting slices
	myslice := []int{45, 67, 90, 33, 21, 56, 78, 99, 15}
	sort.Ints(myslice)
	fmt.Println(myslice)
}
