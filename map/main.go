package main

import "fmt"

func main() {
	statuses := make(map[string]int)

	// append values to map:
	statuses["pending"] = 0
	statuses["inited"] = 1
	statuses["running"] = 2
	statuses["timedout"] = 3
	statuses["successful"] = 4
	statuses["failed"] = 5

	// Reading data from map
	var successfulStatus = statuses["successful"]
	fmt.Println(successfulStatus)

	// Deleting data from map
	delete(statuses, "timedout")
	fmt.Println(statuses)
}
