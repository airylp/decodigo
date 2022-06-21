package main

//decodigo.com

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type Usuario struct {
	XMLName  xml.Name `xml:"usuario"`
	Id       int      `xml:"id,attr"`
	Nombre   string   `xml:"nombre"`
	Username string   `xml:"username"`
	Password string   `xml:"password"`
	Email    string   `xml:"email"`
}

func main() {
	usuario := &Usuario{Id: 1, Nombre: "Jose", Username: "jose", Password: "123", Email: "correo@prueba.com"}

	out, _ := xml.MarshalIndent(usuario, " ", "  ")
	fmt.Println(string(out))

	f, err := os.Create("archivo.xml")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(string(out))

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("listo")
}
