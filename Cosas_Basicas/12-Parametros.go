package main

import "fmt"

// podemos crear una funcion y dentro de una funcion otra funcion esto se llama clousures
func main() {
	pantalon("rojo", "Largo", "Sin bolsillos", "Nike")
}

func pantalon(caracteristicas ...string) {
	//for each
	for _, caracteristica := range caracteristicas {
		fmt.Println(caracteristica)
	}
}
