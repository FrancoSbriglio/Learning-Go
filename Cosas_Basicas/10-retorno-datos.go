package main

import "fmt"

func main() {
	fmt.Println(devolverTexto())
}

func devolverTexto() (string, string) {

	dato1 := "Victor"
	dato2 := "Apellido"

	return dato1, dato2
}
