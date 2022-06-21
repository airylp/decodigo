package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//decodigo.com

type Usuario struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func main() {
	db, err := sql.Open("mysql", "dec01:dec01@tcp(localhost:3306)/pruebas?parseTime=true")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()
	results, err := db.Query("SELECT id, username, password, email FROM usuarios")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var usuario Usuario
		err = results.Scan(&usuario.ID, &usuario.Username, &usuario.Password, &usuario.Email)
		if err != nil {
			panic(err.Error())
		}
		log.Printf(usuario.Username + " " + usuario.Password + " " + usuario.Email)
	}
}
