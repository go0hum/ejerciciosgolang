package main

import (
	"fmt"
	"strconv"
)

var numero int
var texto string
var status bool = true

func main() {
	var numero2, numero3, numero4 int
	numero2, numero3, numero4, texto, texto2, status := 2, 5, 67, "Este es mi texto", "Otro texto", false
	//numero3 := 4
	//numero3 = 15

	var moneda int64 = 0

	numero2 = int(moneda)
	texto = fmt.Sprintf("%d", moneda)

	texto2 = strconv.Itoa(int(moneda))

	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(texto)
	fmt.Println(texto2)
	fmt.Println(status)

	MostrarStatus()
}

//Inicio de mayusculas es para que los Scope sean visibles en otros paquetes
func MostrarStatus() { 
	fmt.Println(status)
}