package main

import "fmt"

//go es tipado, siempre necesitamos definir un tipo de dato a cada una de las variables

func main() {
	var numero1 int = 10
	var numero2 int = 6

	//suma
	fmt.Print("La suma es: ")
	fmt.Println(numero1 + numero2)

	//resta
	fmt.Print("La resta es: ")
	fmt.Println(numero1 - numero2)

	//multiplicacion
	fmt.Print("La multiplicacion es: ")
	fmt.Println(numero1 * numero2)

	//division
	fmt.Print("La division es: ")
	fmt.Println(numero1 / numero2)

	//resto
	fmt.Print("El resto de la division es: ")
	fmt.Println(numero1 % numero2)
}
