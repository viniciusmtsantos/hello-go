package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body) // 	leitura para entrada e saida de dados
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario usuario                                                // criando usuario do tipo usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil { // unMarshal para converter de JSON para struct sendo o primeiro para indicar a requisição em JSON e o endereço de memoria onde ela vai jogar estes dados (com & para indicar o endereco)
		w.Write([]byte("Erro ao converter o usuario para struct"))
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close()
	// PREPARE STATEMENT para evitar um ataque chamado sql injection
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?,?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close() // precisa ser fechado

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email) // substitui os pontos de interrogação
	if erro != nil {
		w.Write([]byte("Erro o executar o statement!"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido"))
		return
	}

	// status code abaixo
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso! Id: %d", idInserido)))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	// SELECT * FROM USUARIOS

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar os usuários"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario // slice de usuario

	// eu vou ler os dados que estão na linha abaixo, jogar no struct usuario e depois jogar este usuario no slice de usuarios
	for linhas.Next() {
		var usuario usuario // usuario do tipo usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil { // scan funciona como se tivesse escanenando e jogando nas propriedades do usuario
			w.Write([]byte("Erro ao escanear o usuario"))
			return
		}
		usuarios = append(usuarios, usuario) // joga o nosso usuario no nosso slice de usuarios
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil { // NewEncoder(w) esta codificando os nossos dados para json e os dados são passados pelo Encode com o slice de usuarios
		w.Write([]byte("Erro ao converter os usuários para JSON"))
		return
	}
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r) // função que retorna os parametros da requisição para nós -> usuarios/{id} pegando o {id}
	// o parametro vem como string e ele é um map então tem que mudar pra inteiro da seguinte forma
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32) // base do numero e o numero de bits
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}
	// e depois eu faço a conexão com o banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar o usuario"))
		return
	}

	var usuario usuario

	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			return
		}
	}

	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao scanear o usuario"))
		return
	}
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Falha ao converter o parâmetro para inteiro"))
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuario para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}

	defer statement.Close()
	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o usuario!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Falha ao converter o parâmetro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete * from usuarios where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao deletar o usuario especificado."))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar o usuario!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
