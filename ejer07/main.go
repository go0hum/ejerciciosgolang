package main

import "fmt"

var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Suma 5 + 7 = %d \n", Calculo(5, 7))

	//Restamos
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}

	fmt.Printf("Restamos 6 - 4 = %d \n", Calculo(6, 4))

	//Dividmos
	Calculo = func(num1 int, num2 int) int {
		return num1 / num2
	}

	fmt.Printf("Dividimos 12 / 3 = %d \n", Calculo(12, 3))
}

func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}
	fmt.Println(resultado())
}

//Closure
