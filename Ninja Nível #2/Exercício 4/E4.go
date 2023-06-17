package main

import "fmt"

// Crie um programa que:
// 		Atribua um valor int a uma variável
// 		Demonstre este valor em decimal, binário e hexadecimal
// 		Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
// 		Demonstre esta outra variável em decimal, binário e hexadecimal

func main() {
	x := 100
	fmt.Printf("%d\t, %b\t, %#x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d\t, %b\t, %#x\t", y, y, y)
}
