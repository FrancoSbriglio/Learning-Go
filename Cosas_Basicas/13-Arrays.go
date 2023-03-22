package main

import "fmt"

// podemos crear una funcion y dentro de una funcion otra funcion esto se llama clousures
func main() {

	//Una de las maneras de definir un array
	/*
		var peliculas [3]string

		peliculas[0] = "Franco"
		peliculas[1] = "Exequiel"
		peliculas[2] = "Sbriglio"

		fmt.Println(peliculas)
		fmt.Println(peliculas[0])
		fmt.Println(peliculas[1])
		fmt.Println(peliculas[2])
	*/
	//otra manera de definir un array

	peliculas := [3]string{"Franco", "Exequiel",
		"Sbriglio"}

	fmt.Println(peliculas)

}
