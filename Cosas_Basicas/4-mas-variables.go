package main

import (
	"fmt"
)

func main() {
	var suma int = 8 + 9
	var resta int = 6 - 4
	var nombre string = "Franco"

	var apellidos string = "Sbriglio"
	apellidos = "Sanfilippo"

	pais := "Argentina" //Operador corto

	fmt.Println("Hola mundo desde Go con" + nombre + apellidos + pais)
	fmt.Println(suma)
	fmt.Println(resta)
}
