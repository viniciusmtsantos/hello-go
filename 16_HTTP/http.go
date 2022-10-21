package main

import (
	"log"
	"net/http"
)

/*
HTTP É UM PROTOCOLO DE COMUNICAÇÃO WEB

CLIENTE (FAZ REQUISIÇÃO) - SERVIDOR (PROCESSA REQUISIÇÃO E ENVIA RESPOSTA)
Request - Response

Rotas:
URI Identificador do Recurso
Método - GET, POST, PUT, DELETE

Uae eC ee)
a ea ee ats
*/

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar paginas de Usuarios"))
}

func main() {
	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))

}
