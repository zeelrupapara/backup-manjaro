package main

import (
	"fmt"
)

// for reset the stack or create empty stack
func stack_empty(stack []int) []int {
	stack = nil
	fmt.Println("Reset Stack")
	fmt.Println(stack)
	return stack
}

// for push a item in to the stack
func push(stack []int, item int) []int {
	stack = append(stack, item)
	fmt.Println("Pushe item in the stack is", item)
	fmt.Println(stack)
	return stack
}

// for remove last added item in the stack
func pop(stack []int) []int {
	// check is stack is empty or not?
	if len(stack) == 0 {
		fmt.Println("item pop in the queue")
		fmt.Println(stack)
		return stack
	} else {
		stack = stack[:len(stack)-1]
		fmt.Println("item pop in the queue")
		fmt.Println(stack)
		return stack
	}
}

func main() {
	var stack []int = nil
	stack = push(stack, 1)
	stack = push(stack, 2)
	stack = push(stack, 4)
	stack = push(stack, 10)
	stack = pop(stack)
	stack = stack_empty(stack)
	stack = push(stack, 11)
	stack = push(stack, 12)
	stack = push(stack, 13)
	pop(stack)
}
