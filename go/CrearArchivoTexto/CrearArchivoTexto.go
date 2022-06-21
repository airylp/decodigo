package main

//decodigo.com

import (
	"fmt"
	"os"
)

var path = "prueba.txt"

func main() {
	crearArchivo()
	escribeArchivo()
}

func crearArchivo() {
	//Verifica que el archivo existe
	var _, err = os.Stat(path)

	//Crea el archivo si no existe
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if existeError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("Archivo creado exitosamente", path)
}

func escribeArchivo() {
	// Abre archivo usando permisos READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if existeError(err) {
		return
	}
	defer file.Close()

	// Escribe algo de texto linea por linea
	_, err = file.WriteString("Hola \n")
	if existeError(err) {
		return
	}
	_, err = file.WriteString("Mundo \n")
	if existeError(err) {
		return
	}

	// Salva los cambios
	err = file.Sync()
	if existeError(err) {
		return
	}

	fmt.Println("Archivo actualizado existosamente.")
}

func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
