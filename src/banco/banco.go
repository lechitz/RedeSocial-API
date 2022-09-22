package banco

import (
	"api/src/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //Driver
)

//Conectar abre a conexão com o Banco de Dados e a retorna
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close() //O Banco de Dados só é fechado caso não consiga fazer mais o Ping
		return nil, erro
	}

	return db, nil
}