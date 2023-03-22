package main

import "fmt"

// podemos crear una funcion y dentro de una funcion otra funcion esto se llama clousures
func main() {

	//Una de las maneras de definir un array

	var peliculas [3][2]string

	peliculas[0][0] = "Franco"
	peliculas[0][1] = "Exequiel"

	peliculas[1][0] = "Hola"
	peliculas[1][1] = "Mundo"

	peliculas[2][0] = "GO"
	peliculas[2][1] = "NetCore"

	fmt.Println(peliculas)

	fmt.Println(peliculas[1][0])
}
