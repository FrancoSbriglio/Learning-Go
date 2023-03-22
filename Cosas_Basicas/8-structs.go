package main

import "fmt"

type Gorra struct { // tipos de datos personalizados
	marca  string
	color  string
	precio float32
	plana  bool
}

func main() {

	var gorra_negra = Gorra{
		marca:  "Nike",
		color:  "Negro",
		precio: 25.50,
		plana:  false}

	var gorra_negra_1 = Gorra{
		"Nike",
		"Negro",
		25.50,
		false}

	fmt.Println(gorra_negra)
	fmt.Println(gorra_negra_1)
	fmt.Println(gorra_negra_1.marca)
	fmt.Println(gorra_negra_1.precio)
}
