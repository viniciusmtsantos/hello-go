package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Conectar abre a conexao com o banco de dados
func Conectar() (*sql.DB, error) { // retornando ponteiro de sql e um erro

	stringConexao := "root:teste@tcp(localhost:3307)/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao) // mutualmente exclusivos um não existe com o outro
	if erro != nil {
		return nil, erro // retornando o valor zero e o erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
