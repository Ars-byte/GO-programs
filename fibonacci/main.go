package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Ingrese un número: ")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Entrada no válida")
		return
	}

	if n < 0 {
		fmt.Println("El número debe ser mayor o igual a cero")
		return
	}

	fibonacci := make([]int, n+1)
	fibonacci[0] = 0
	if n > 0 {
		fibonacci[1] = 1
	}

	for i := 2; i <= n; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}

	fmt.Println("Secuencia de Fibonacci hasta", n, ":")
	for i := 0; i <= n; i++ {
		fmt.Print(fibonacci[i], " ")
	}
	fmt.Println()
}
