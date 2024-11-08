package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// initializePostgres configura a conexão com o banco de dados Postgres e retorna a instância *sql.DB.
func initializePostgres() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=hepherius password=2474 dbname=gastos sslmode=disable")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
