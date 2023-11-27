package main

import "fmt"

func main() {
	queue := NewQueue()
	queue.Enqueue(100)
	fmt.Println("queue:", queue)

	queue.Enqueue(25).Enqueue(77)
	fmt.Println("queue:", queue)
	fmt.Println("Is queue empty?", queue.IsEmpty())

	result, err := queue.Peek()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The next item in the queue:", result)
	}

	fmt.Println("Dequeue one item...")
	result, err = queue.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	fmt.Println("Dequeue one item...")
	result, err = queue.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	fmt.Println("Dequeue one item...")
	result, err = queue.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

// Queue struct
type Queue struct {
	data []int
}

// NewQueue returns a new Queue
func NewQueue() *Queue {
	return &Queue{
		data: []int{},
	}
}

// IsEmpty returns true if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

// Peek returns the next element in the queue without removing it
func (q *Queue) Peek() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Queue is empty")
	}
	return q.data[0], nil
}

// Enqueue adds an element to the end of the queue and returns a pointer to it
func (q *Queue) Enqueue(n int) *Queue {
	q.data = append(q.data, n)
	return q
}

// Dequeue removes and returns the front element from the queue
func (q *Queue) Dequeue() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Queue is empty")
	}

	element := q.data[0]
	q.data = q.data[1:]
	return element, nil
}
