package main

import "fmt"

func operators3() {
	var a int = 14
	var b int = 30

	// Logical AND operator (&&)
	if a != b && a <= b {
		fmt.Println("true")
	}

	// Logical OR operator (||)
	if a != b || a <= b {
		fmt.Println("true")
	}

	// Logical NOT operator (!)
	if !(a == b) {
		fmt.Println("true")
	}
}
