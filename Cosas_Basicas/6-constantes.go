package main

import (
	"fmt"
)

//go es tipado, siempre necesitamos definir un tipo de dato a cada una de las variables

func main() {
	var suma int = 8 + 9
	var resta int = 6 - 4
	var nombre string = "Franco"

	var apellidos string = "Sbriglio"
	apellidos = "Sanfilippo"

	pais := "Argentina" //Operador corto

	var prueba bool = true

	var prueba1 float32 = 1.324

	const year int = 2023

	fmt.Println(year)

	fmt.Println(prueba)
	fmt.Println(prueba1)
	fmt.Println("Hola mundo desde Go con" + nombre + apellidos + pais)
	fmt.Println(suma)
	fmt.Println(resta)
}
