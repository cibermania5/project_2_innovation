/*
	Author:  Mike Motta
	Purpose: This is the main file
*/

package main

import (
	"./operations"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func main() {
	r:= mux.NewRouter()
	r.Use(Recovery)



	r.HandleFunc("/casas/consulta/{id}", operations.GetCasa).Methods("GET")
	r.HandleFunc("/casas/nuevo", operations.CreaCasa).Methods("POST")
	r.HandleFunc("/casas/multas/{id}", operations.AniadeMultas).Methods("PATCH")
	r.HandleFunc("/casas/debe/{id}", operations.CambiaDebe).Methods("PATCH")
	r.HandleFunc("/casas/consulta", operations.GetTodos).Methods("GET")
	r.HandleFunc("/casas/saldo/{id}", operations.CalculaTotalCasa).Methods("GET")
	r.HandleFunc("/casas/pagar/{id}", operations.Pagar).Methods("PATCH")
	r.HandleFunc("/mensajes/nuevo", operations.CreaMensaje).Methods("POST")
	r.HandleFunc("/mensajes/consulta/{id}", operations.GetMensaje).Methods("GET")
	r.HandleFunc("/mensajes/consulta", operations.GetMensajes).Methods("GET")
	r.HandleFunc("/mensajes/actualiza/{id}", operations.ActualizaMensajes).Methods("PATCH")


	//log.Fatal("falló por: ", http.ListenAndServe(":5801", r))
	err := http.ListenAndServe(":5801", r)
	if err != nil {
		fmt.Println("falló por: ", err)
		log.Fatal("falló por: ", err)
	}

}


// https://medium.com/@masnun/panic-recovery-middleware-for-go-http-handlers-51147c941f9
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				fmt.Println(err) // May be log this error? Send to sentry?

				jsonBody, _ := json.Marshal(map[string]string{
					"error": "Hubo un problema",
				})

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonBody)
			}

		}()

		next.ServeHTTP(w, r)

	})
}
