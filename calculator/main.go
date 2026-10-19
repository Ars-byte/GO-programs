package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func runCalculator() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese un número:")
	if !scanner.Scan() {
		return
	}
	num1, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println("Número no válido")
		return
	}

	fmt.Println("Ingrese otro número:")
	if !scanner.Scan() {
		return
	}
	num2, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println("Número no válido")
		return
	}

	fmt.Println("Ingrese una operación: +, -, *, /")
	if !scanner.Scan() {
		return
	}
	operacion := scanner.Text()

	fmt.Println("Resultado:")
	switch operacion {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		if num2 == 0 {
			fmt.Println("No se puede dividir entre cero")
			return
		}
		fmt.Println(num1 / num2)
	default:
		fmt.Println("Operación no válida")
	}
}

func main() {
	runCalculator()
}
