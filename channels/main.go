package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func main() {

// 	channel := make(chan string)
// 	go func() {
// 		channel <- "banana"
// 		channel <- "apple"
// 	}()

// 	fmt.Println(<-channel)
// 	fmt.Println(<-channel)

// }

func main() {
	channel := make(chan int)
	go getRandomNumber(channel)
	for randomNumber := range channel {
		fmt.Println("Radnom number is:", randomNumber)
	}
}

func getRandomNumber(channel chan int) {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < 3; i++ {
		number := rand.Int(1000)
		time.Sleep(time.Second * 1)
		channel <- number
	}
	close(channel)
}
