package main

import (
	"fmt"
	"os"
)

// main lee un archivo de texto y muestra su contenido en la terminal.
//
// Uso:
//
//	go run .                -> lee ./hello.txt (archivo de ejemplo)
//	go run . otro.txt       -> lee el archivo que indiques
//
// El programa debe ejecutarse dentro de esta carpeta (read-files), porque la
// ruta por defecto es relativa al directorio actual.
func main() {
	// Si el usuario pasa una ruta como argumento la usamos; si no, hello.txt.
	path := "hello.txt"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	// os.ReadFile lee el archivo completo y devuelve su contenido como []byte.
	data, err := os.ReadFile(path)
	if err != nil {
		// Si el archivo no existe o no tenemos permiso, avisamos y salimos
		// con un código de error en lugar de imprimir contenido vacío.
		fmt.Fprintf(os.Stderr, "no se pudo leer %q: %v\n", path, err)
		os.Exit(1)
	}

	// Los bytes leídos son texto, así que los convertimos a string para
	// imprimirlos tal cual estaban escritos en el archivo.
	fmt.Print(string(data))
}
