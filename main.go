/*
	Author:  Mike Motta
	Purpose: This is the main file
*/

package main

import (
	"./operations"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func main() {
	r:= mux.NewRouter()

	r.HandleFunc("/casas/consulta/{id}", operations.GetCasa).Methods("GET")
	r.HandleFunc("/casas/nuevo", operations.CreaCasa).Methods("POST")
	r.HandleFunc("/casas/multas/{id}", operations.AniadeMultas).Methods("PATCH")
	r.HandleFunc("/casas/consulta", operations.GetTodos).Methods("GET")
	r.HandleFunc("/casas/saldo/{id}", operations.CalculaTotalCasa).Methods("GET")
	r.HandleFunc("/casas/pagar/{id}", operations.Pagar).Methods("PATCH")
	r.HandleFunc("/mensajes/nuevo", operations.CreaMensaje).Methods("POST")
	r.HandleFunc("/mensajes/consulta/{id}", operations.GetMensaje).Methods("GET")
	r.HandleFunc("/mensajes/consulta", operations.GetMensajes).Methods("GET")
	r.HandleFunc("/mensajes/actualiza/{id}", operations.ActualizaMensajes).Methods("PATCH")
	log.Fatal(http.ListenAndServe(":5801", r))

}
