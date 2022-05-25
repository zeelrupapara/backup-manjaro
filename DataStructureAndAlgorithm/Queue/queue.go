package main

import (
	"fmt"
)

// for reset the queue or create empty queue
func empty_queue(queue []int) []int {
	queue = nil
	fmt.Println("Reset queue")
	fmt.Println(queue)
	return queue
}

// for push a item in to the queue
func push(queue []int, item int) []int {
	queue = append(queue, item)
	fmt.Println("Pushed item in the queue is", item)
	fmt.Println(queue)
	return queue
}

// for remove first added item in the queue
func pop(queue []int) []int {
	// check is queue is empty or not?
	if len(queue) == 0 {
		fmt.Println("item pop in the queue")
		fmt.Println(queue)
		return queue
	} else {
		queue = queue[1:]
		fmt.Println("item pop in the queue")
		fmt.Println(queue)
		return queue
	}
}

func main() {
	var queue []int = nil
	queue = push(queue, 1)
	queue = push(queue, 2)
	queue = push(queue, 4)
	queue = push(queue, 10)
	queue = pop(queue)
	queue = empty_queue(queue)
	queue = push(queue, 11)
	queue = push(queue, 12)
	queue = push(queue, 13)
	pop(queue)
}
