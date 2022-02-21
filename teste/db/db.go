package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"squirrel/teste/domain"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *sql.DB
var err error

func Connection() {

	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbname := os.Getenv("NAME")
	dbpassword := os.Getenv("PASSWORD")

	// Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname, dbpassword, dbPort)

	DB, err = sql.Open(dialect, dbURI)

	if err != nil {
		panic("Conexão mal sucedida  :(")
	} else {
		fmt.Println("Conexão ok")
	}
	//defer DB.Close()
}

func FindAll() []domain.Cliente {

	clientes := sq.Select("id, name,email,fone,tipo_cliente").From("clientes")

	rows, sqlerr := clientes.RunWith(DB).Query()
	if sqlerr != nil {
		panic(fmt.Sprintf("QueryRow failed: %v", sqlerr))
	}

	var listaClientes []domain.Cliente
	for rows.Next() {
		var cliente domain.Cliente
		if err := rows.Scan(&cliente.ID, &cliente.Name, &cliente.Email, &cliente.Fone, &cliente.TipoCliente); err != nil {
			log.Fatal(err)
		}

		listaClientes = append(listaClientes, cliente)
	}
	fmt.Println(listaClientes)
	return listaClientes
}

func Create() {
	clienteTeste := domain.Cliente{}
	clienteTeste.ID = 7
	clienteTeste.Name = "Teste"
	clienteTeste.Email = "teste@testeclienteTeste"
	clienteTeste.Fone = 996282879
	clienteTeste.TipoCliente = 3

	sq.Insert("clientes").Columns("id, name,email,fone,tipo_cliente").Values(clienteTeste.ID, clienteTeste.Name, clienteTeste.Email, clienteTeste.Fone, clienteTeste.TipoCliente)
}
