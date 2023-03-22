package main

import "fmt"

type Gorra struct { // tipos de datos personalizados
	marca  string
	color  string
	precio float32
	plana  bool
}

func main() {
	var numero1 float32 = 10
	var numero2 float32 = 6
	holaMundo()
	//suma
	fmt.Print("La suma es: ")
	fmt.Println(operacion(numero1, numero2, "+"))

	//resta
	fmt.Print("La resta es: ")
	fmt.Println(operacion(numero1, numero2, "-"))

	//multiplicacion
	fmt.Print("La multiplicacion es: ")
	fmt.Println(operacion(numero1, numero2, "*"))

	//division
	fmt.Print("La division es: ")
	fmt.Println(operacion(numero1, numero2, "/"))

}

func holaMundo() {
	fmt.Println("Hola mundo")
}

func operacion(n1 float32, n2 float32, op string) float32 {

	var resultado float32

	if op == "+" {
		resultado = n1 + n2
	}
	if op == "-" {
		resultado = n1 - n2
	}
	if op == "*" {
		resultado = n1 + n2
	}
	if op == "/" {
		resultado = n1 / n2
	}
	return resultado

}
