package main

//decodigo.com

import (
	"encoding/json"
	"fmt"
)

type Usuario struct {
	Nombre   string `json:"nombre"`
	Usuario  string `json:"usuario"`
	Password string `json:"password"`
	Email    string `json:"email"`
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
