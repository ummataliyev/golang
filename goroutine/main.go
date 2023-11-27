package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		display("Hello!")
		wg.Done()
	}()
	go func() {
		go display("Hi!")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Completed!")
}

func display(input string) {
	for i := 1; i <= 7; i++ {
		fmt.Println(i, input)
		time.Sleep(1 * time.Second)
	}
}
