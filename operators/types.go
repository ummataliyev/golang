package main

import "fmt"

func types() {
	// integer
	var X uint8 = 255
	fmt.Println(X+1, X)
	var Y int16 = 32765
	fmt.Println(Y - 2)

	// float
	a := 20.45
	b := 34.89
	c := b - a
	fmt.Printf("Result: %f", c)
	fmt.Printf("\n type of c is: %T", c)

	// complex
	var d1 complex128 = complex(6, 2)
	var d2 complex64 = complex(9, 2)
	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Printf(" type of d1 is %T and type of d2 is %T", d1, d2)

	// boolean
	str1 := "Virtual Lessons"
	str2 := "Virtual Lessons"
	str3 := "Virtual Lessons"
	result1 := str1 == str2
	result2 := str1 == str3
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Printf("type of result1 is %T and type of result2 is %T", result1, result2)

	// string
	str := "Virtual Lessons"
	fmt.Printf("\nlength of string is %d", len(str))
	fmt.Printf("\nString: %s", str)
	fmt.Printf("\ntype of string is %T", str)
}
