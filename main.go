package main

import "fmt"

//складываем
func sum(a, b int) int {
	return a + b
}

//вычитаем
func min(a, b int) int {
	return a - b
}

func main() {
	fmt.Println("Привет")
	fmt.Println("Это калькулятор")
	a := 5
	b := 8
	fmt.Printf("Давай сложим %d и %d. Мы получаем %d\n", a, b, sum(a, b))
	fmt.Printf("Давай вычтем %d и %d. Мы получаем %d\n", a, b, min(a, b))
}
