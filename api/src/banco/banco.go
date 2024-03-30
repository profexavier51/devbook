package banco

import (
	"api/src/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o MySQL
)

func Conectar() (*sql.DB, error) {

	db, erro := sql.Open("mysql", config.StringConexao)
    if erro != nil {
        return nil, fmt.Errorf("erro ao conectar ao banco de dados: %v", erro)
    }

    // Testa a conexão com o banco de dados
    if erro = db.Ping(); erro != nil {
        db.Close() // Fecha a conexão em caso de erro
        return nil, fmt.Errorf("erro ao testar a conexão com o banco de dados: %v", erro)
    }

    fmt.Println("Conexão ao banco de dados MySQL bem-sucedida!")

    return db, nil
}