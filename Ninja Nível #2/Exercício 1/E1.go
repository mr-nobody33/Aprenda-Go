package main

// Escreva um programa que mostre um número em decimal, binário, e hexadecimal.

import "fmt"

func main() {
	x := 10
	fmt.Printf("%d, %#x, %b", x, x, x)

	// %d = decimal
	// %#x = hexadecimal
	// %b = binário
}
