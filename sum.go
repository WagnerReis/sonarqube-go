package main

import "fmt"

func main() {
	fmt.Printf("Sum: %d\n", sum(1, 2))
}

func sum(a int, b int) int {
	return a + b
}