package main

import (
	"log"
	"net/http"
)

/*
HTTP É UM PROTOCOLO DE COMUNICAÇÃO WEB

CLIENTE (FAZ REQUISIÇÃO) - SERVIDOR (PROCESSA REQUISIÇÃO E ENVIA RESPOSTA)
Request - Response

ROTAS:
URI - Identificador do recurso (mensagem que está sendo enviada)
Método - GET, POST, PUT, DELETE

*/

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) { // dois parametros especificos que vao fazer o processo de request e response
	w.Write([]byte("Carregar paginas de Usuarios")) // esse cara escreve uma mensagem com o tipo de slice de byte
}

func main() {

	// definimos a URI e o metodo que ela irá receber
	http.HandleFunc("/home", home) // o primeiro é o URI da rota e o segundo é a função (tem uma assinatura especifica para ser reconhecida) que vai receber a requisição e vai saber como lidar com ela

	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil)) // É nesta porta que vamos ficar ouvindo as requisições e dando as respostas e o segundo parâmetro é nil por enquanto

	// sem a rota não da pra fazer nenhuma requisição

}
