package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/* CREATE - POST
READ - GET
UPTADE - PUT
DELETE - DELETE
*/

func main() {
	router := mux.NewRouter() // vai conter todas as rotas para fazer a interação com o banco de dados

	// usamos o pacote http para subir o servidor junto com o router
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost) // ou "POST"
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", servidor.DeletarUsuario).Methods(http.MethodDelete)

	fmt.Println("Escutando na 5000")
	log.Fatal(http.ListenAndServe(":5000", router)) // passamos o router e não o nil como no caso da aula de HTTP

}
