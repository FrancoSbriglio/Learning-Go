package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //w response writer, r request
		fmt.Fprintf(w, "Hola mundo desde mi servidor web con go")
	}) //Creando el servidor, esto nos sirve para crear una ruta, que va a ser la ruta principal de la aplicacion o del servidor

	fmt.Println("El servidor deberia estar corriendo en http://localhost:8080")
	server := http.ListenAndServe(":8080", nil) //escucha y sirve para que el servidor se levante indicando el puerto

	log.Fatal(server)

}
