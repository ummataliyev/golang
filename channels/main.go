package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan int)
	go getRandomNumber(channel)
	for randomNumber := range channel {
		fmt.Println("Random number is:", randomNumber)
	}
}

func getRandomNumber(channel chan int) {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 3; i++ {
		randomNumber := rand.Intn(100) // Generate a random number between 0 and 99
		channel <- randomNumber
		time.Sleep(time.Second) // Introduce a delay to demonstrate receiving in the main goroutine
	}
	close(channel)
}
