package main

import "fmt"

//- Utilizando a solução do exercício anterior:
//1. Em package-level scope, utilizando a palavra-chave var, crie uma variável com o identificador "y". O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.
//2. Na função main:
//		1. Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, utilizando o operador "=", atribua o valor de "x" a "y"
//		2. Demonstre o valor de "y"
//		3. Demonstre o tipo de "y"

type hotdog int

var x hotdog

var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	fmt.Println("↑ foi x. \n↓ é y!")
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
