// Bitwise Operators

package main

import "fmt"

func operators4() {
	a := 14
	b := 30

	// & (bitwise AND)
	result1 := a & b
	fmt.Printf("Result of a & b is %d", result1)

	// | (bitwise OR)
	result2 := a | b
	fmt.Printf("Result of a | b is %d", result2)

	// ^ (bitwise XOR)
	result3 := a ^ b
	fmt.Printf("Result of a ^ b is %d", result3)

	// << (left shift)
	result4 := a << 1
	fmt.Printf("Result of p is %d", result4)
}
