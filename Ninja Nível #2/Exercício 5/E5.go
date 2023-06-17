package main

import "fmt"

// Crie uma variável de tipo string utilizando uma raw string literal.
// Demonstre-a.

func main() {
	x := `Uma 
					coisa
								é
				qualquer
										coisa.`
	fmt.Println(x)
}
