package main

import "fmt"

// The first way to using for loop

// func main() {
// 	for i := 1; i < 8; i++ {
// 		fmt.Println(i)
// 	}
// }

// The second way to use for loop

// func main() {
// 	for {
// 		fmt.Println("Starting " + strconv.Itoa(time.Now().Second()))
// 		time.Sleep(1 * time.Second)
// 	}

// }

// The third way to use for loop

// func main() {
// 	i := 1
// 	for i < 10 {
// 		fmt.Println("Waiting for " + strconv.Itoa(i))
// 		i++
// 	}
// }

// The fourth way to use for loop with lists
// func main() {
// 	myarray := [3]string{"earth", "sun", "moon"}
// 	for index, value := range myarray {
// 		fmt.Println("index:", index, "value:", value)
// 	}
// }

// The fifth way to use for lopp with map (dict)
func main() {
	myarray := map[int]string{
		1: "Python",
		2: "Golang",
		3: "React",
	}

	for key, value := range myarray {
		fmt.Println("Key:", key, "Value:", value)
	}
}
