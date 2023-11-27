package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		for {
			channel1 <- "Speeed"
			time.Sleep(time.Hour * 1)
		}
	}()

	go func() {
		for {
			channel2 <- "Low"
			time.Sleep(time.Minute * 30)
		}
	}()

	for {
		select {
		case message1 := <-channel1:
			fmt.Println(message1)

		case message2 := <-channel2:
			fmt.Println(message2)

		default:
			fmt.Println("Unknown")
		}
	}
}
