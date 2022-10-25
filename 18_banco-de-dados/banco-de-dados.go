package main

import (
	"database/sql" // quem usa o pacote do mysql abaixo é este aí
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // importe implícito deste driver ali no pacote database/sql
)

func main() {

	// string de conexão passando usuario e a senha usuario:senha@/banco e depois do ? tem parametros adicionais
	stringConexao := "root:teste@tcp(localhost:3307)/devbook?charset=utf8&parseTime=True&loc=Local" // charset para carac latinos e parseTime para dar o parse se tiver campo de tempo

	// agora eu chamo um metodo sql.open para chamar a conexao pelo pacote padrao do golang
	// Este metodo retorna dois valores, o primeiro uma conexao e o segundo um erro

	db, erro := sql.Open("mysql", stringConexao) // mysql é o tipo de banco que estou me referindo e o outro é o string de conexão
	if erro != nil {                             // tratando o erro
		fmt.Println("dentro do open")
		log.Fatal(erro)
	}
	defer db.Close() // ele fecha a conexão indepedente do que aconteça antes de o main terminar

	if erro = db.Ping(); erro != nil { // se ele não cair dentro deste ping, a conexão estará funcionando normal
		fmt.Println("dentro do ping")
		log.Fatal(erro)
	}

	fmt.Println("Conexao esta aberta")

	linhas, erro := db.Query("select * from usuarios") // muitas funcoes retornam erro tbm, como neste caso

	if erro != nil {
		log.Fatal(erro)
	}

	defer linhas.Close() // podemos fechar as linhas, pois as linhas servem como cursor, e não precisam ficar abertos

	fmt.Println(linhas)
}
