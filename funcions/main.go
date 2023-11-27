package main

import "fmt"

// func main() {
// 	// result := sum(1, 2)
// 	result, err := sqrt(144)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("Result of sqrt:", result)
// 	}
// }

// func sum(x int, y int) int {
// 	return x + y
// }

// func sqrt(x float64) (float64, error) {
// 	if x < 0 {
// 		return 0, errors.New("sqrt must be positive")
// 	}
// 	return math.Sqrt(x), nil
// }

// By value
// func main() {
// 	var a int = 10
// 	var b int = 20
// 	fmt.Println(a, b)
// 	swap(a, b)
// 	fmt.Println(a, b)
// }

// func swap(x, y int) int {
// 	var o int
// 	o = x
// 	x = y
// 	y = o
// 	return 0
// }

// By reference
func main() {
	var a int = 10
	var b int = 20
	fmt.Println(a, b)
	swapByRef(&a, &b)
	fmt.Println(a, b)
}

func swapByRef(x, y *int) {
	var o int
	o = *x
	*x = *y
	*y = o
}
