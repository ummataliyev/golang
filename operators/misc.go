// Misc Operators

package main

import "fmt"

func operators6() {
	a := 4

	// using adress of cache (&) and pointer operators
	b := &a
	fmt.Println(*b)
	*b = 7
	fmt.Println(a)
}
