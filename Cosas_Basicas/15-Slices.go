package main

import "fmt"

// arrays dinamicos - slices
func main() {

	peliculas := []string{"Franco", "Exequiel",
		"Sbriglio"}

	peliculas = append(peliculas, "Sanfilippo")
	peliculas = append(peliculas, "Pepe")
	fmt.Println(len(peliculas)) // longitud
	fmt.Println(peliculas[0:3]) // pedir mas de un objeto
	fmt.Println(peliculas[4])

}
