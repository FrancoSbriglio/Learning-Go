package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/peliculas", MovieList)
	router.HandleFunc("/peliculas/{id}", MovieShow)
	server := http.ListenAndServe(":8080", router) //escucha y sirve para que el servidor se levante indicando el puerto

	log.Fatal(server)

}
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

func MovieList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Listado de peliculas")
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	movie_id := params["id"]

	fmt.Fprintf(w, "Has cargo la pelicula numero %s", movie_id)
}
