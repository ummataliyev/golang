// Mathematical Operations

package main

import "fmt"

func operators1() {
	a := 64
	b := 20

	// plus
	result1 := a + b
	fmt.Printf("Result of a + b = %d\n", result1)

	// minus
	result2 := a - b
	fmt.Printf("Result of a - b = %d\n", result2)

	// multiplication
	result3 := a * b
	fmt.Printf("Result of a * b = %d\n", result3)

	// division
	result4 := a / b
	fmt.Printf("Result of a / b = %d\n", result4)

	// remainder (modul)
	result5 := a % b
	fmt.Printf("Result of a %% b = %d\n", result5)
}
