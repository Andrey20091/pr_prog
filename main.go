package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func min(a, b int) int {
	return a - b
}

func main() {
	fmt.Println("Hello")
	fmt.Println("It is calculate app")
	a := 5
	b := 8
	fmt.Printf("Let's sum %d and %d. We have now %d\n", a, b, sum(a, b))
	fmt.Printf("Lets min %d and %d. We have now %d\n", a, b, min(a, b))
}
