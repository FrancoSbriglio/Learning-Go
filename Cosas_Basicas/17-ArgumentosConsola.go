package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("****Mi programa con Go****")

	fmt.Println("Hola " + os.Args[1] + " bienvenido al programa con go")

	edad, _ := strconv.Atoi(os.Args[2])

	if edad >= 18 && edad <= 99 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("No eres mayor de edad")
	}

}
