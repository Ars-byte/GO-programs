package main

import (
	"fmt"
	"time"
)

func main() {
	tiempo_actual := time.Now()
	formateo_de_tiempo := tiempo_actual.Format("2006-01-02 15:04:05")
	fmt.Println("Fecha y hora actual:", formateo_de_tiempo)
}
