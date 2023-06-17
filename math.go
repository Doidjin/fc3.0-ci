package main

import "fmt"

func main() {
	fmt.Println("A soma é:", Soma(1, 2))
}

func Soma(a int, b int) int {
	return a + b
}
