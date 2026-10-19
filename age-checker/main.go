package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese su edad:")
	if !scanner.Scan() {
		return
	}
	edad, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Entrada no válida")
		return
	}

	if edad < 18 {
		fmt.Println("Eres menor de edad")
	} else {
		fmt.Println("Eres mayor de edad")
	}
}
