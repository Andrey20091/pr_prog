package main

import "fmt"

//Умножение
func mult(a, b int) int {
	return a * b
}

func main() {
	fmt.Println("привет")
	fmt.Println("Это калькулятор")
	a := 5
	b := 8
	fmt.Printf("Мы млжем умножать. Имеем %d, %d. После умножения получим %d\n", a, b, mult(a, b))
}
