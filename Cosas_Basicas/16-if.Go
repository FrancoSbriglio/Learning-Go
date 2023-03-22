package main

import "fmt"

func main() {

	fmt.Println("****Mi programa con Go****")

	edad := 18

	if edad >= 18 && edad <= 99 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("No eres mayor de edad")
	}

}
