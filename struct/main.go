package main

import "fmt"

type Car struct {
	Make, Model, Color string
	Year, Weight       int
	Engine             engine
}

type engine struct {
	name string
	hp   int
}

func main() {
	// first way of declaring object
	// var myCar = Car

	// second way of declaring object
	myCar := Car{
		Make:  "BMW",
		Model: "x7",
		Color: "black",
		Year:  2023, Weight: 2525,
		Engine: engine{name: "T8", hp: 400}}

	fmt.Println(myCar)
}
