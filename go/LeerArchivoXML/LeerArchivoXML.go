package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

//decodigo.com

type Usuarios struct {
	Usuarios []Usuario `xml:"usuario"`
}

type Usuario struct {
	Tipo     string `xml:"tipo,attr"`
	Nombre   string `xml:"nombre"`
	Username string `xml:"username"`
	Password string `xml:"password"`
	Correo   string `xml:"correo"`
}

func main() {
	datos, _ := ioutil.ReadFile("xml_ejemplo.xml")

	usuarios := &Usuarios{}

	_ = xml.Unmarshal([]byte(datos), &usuarios)

	fmt.Println("usuarios encontrados:", len(usuarios.Usuarios))
	for i := 0; i < len(usuarios.Usuarios); i++ {
		fmt.Println("tipo: ", usuarios.Usuarios[i].Tipo)
		fmt.Println("nombre: ", usuarios.Usuarios[i].Nombre)
		fmt.Println("username: ", usuarios.Usuarios[i].Username)
		fmt.Println("password: ", usuarios.Usuarios[i].Password)
		fmt.Println("correo: ", usuarios.Usuarios[i].Correo)
	}
}
