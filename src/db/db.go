package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func ConectaBanco() *sql.DB {
	conexao := "user=" + os.Getenv("USER") + " dbname=" + os.Getenv("DB_NAME") + " password=" + os.Getenv("PASS") + " host=" + os.Getenv("HOST") + " port=" + os.Getenv("PORT") + " sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
