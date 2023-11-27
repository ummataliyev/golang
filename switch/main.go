package main

import (
	"fmt"
	"time"
)

func main() {
	testSwitch2()
}

func testIf() {
	i := 7
	if i == 7 {
		fmt.Println("Seven!")
	}
}

// IF, IF ELSE, ELSE

func testIfElse() {
	points := 6000
	if points < 5000 {
		fmt.Println("Silver!")
	} else if points < 10000 {
		fmt.Println("Gold!")
	} else {
		fmt.Println("Platinum!")
	}
}

// Switch cases

func testSwtich1() {
	weekday := time.Now().Weekday()
	switch weekday {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wednesday")
	case 5:
		fmt.Println("Thursday")
	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")
	default:
		fmt.Println("Invalid weekday!")
	}
}

// Another switch cases example
func testSwitch2() {
	greeting := 3
	switch {
	case greeting == 1:
		fmt.Println("Hello!")
	case greeting == 2:
		fmt.Println("Ciao!")
	case greeting == 3:
		fmt.Println("Bonjour!")
	case greeting == 4:
		fmt.Println("Shalom!")
	case greeting == 5:
		fmt.Println("Привет!")
	default:
		fmt.Println("Invalid greeting!")
	}
}
