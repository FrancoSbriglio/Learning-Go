package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	//nuevo_texto := []byte(os.Args[1])
	//escribir := ioutil.WriteFile("fichero.txt", nuevo_texto, 0777) //Crea un archivo nuevo
	//showError(escribir)

	nuevo_texto := os.Args[1] + "\n"
	fichero, err := os.OpenFile("fichero.txt", os.O_APPEND, 0777)
	showError(err)

	escribir, err := fichero.WriteString(nuevo_texto)
	showError(err)
	fmt.Println(escribir)

	fichero.Close()

	texto, errorDeFichero := ioutil.ReadFile("fichero.txt")
	showError(errorDeFichero)

	fmt.Println(string(texto))

}
func showError(e error) {
	if e != nil {
		panic(e)
	}
}
