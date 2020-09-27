package main

import "fmt"

/*var tabla [10]int*/
/*var matriz [5][7]int*/

func main() {
	/*tabla[0] = 1
	tabla[5] = 15

	fmt.Println(tabla)

	tabla := [10]int{5, 6, 98, 1, 4, 5, 6, 8, 9, 10}

	fmt.Println(tabla)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}*/

	/*matriz[3][5] = 1

	fmt.Println(matriz)*/

	/*matriz := []int{2, 4, 5}
	fmt.Println(matriz)*/
	variante2()
	variante3()
	variante4()
}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	//porcion := elementos[3:]
	//porcion := elementos[:4]
	porcion := elementos[2:4]

	fmt.Println(porcion)
}

func variante3() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

}

func variante4() {
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\n Largo %d, Capacidad %d", len(nums), cap(nums))
}
