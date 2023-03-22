package main

import "fmt"

// podemos crear una funcion y dentro de una funcion otra funcion esto se llama clousures
func main() {

	fmt.Print("Pedido 1 ---->")
	fmt.Println(gorras(45, "E"))

	fmt.Print("Pedido 2 ---->")
	fmt.Println(gorras(24, "USD"))
}

func gorras(pedido float32, moneda string) (string, float32, string) {

	precio := func() float32 {
		return pedido * 7
	}

	return "El precio de las gorras pedidas es:", precio(), moneda
}
