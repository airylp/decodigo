package main

import (
	"encoding/json"
	"fmt"
)

type Usuario struct {
	Nombre   string
	Usuario  string
	Password string
	Email    string
}

func main() {
	usuario := &Usuario{Nombre: "Francisco", Usuario: "franz", Password: "123", Email: "correo@prueba.com"}
	usuarioJSON, err := json.Marshal(usuario)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Println(string(usuarioJSON))
}
