package main

import "fmt"

func main() {
	paises := make(map[string]string)
	fmt.Println(paises)
	paises["Mexico"] = "D.F"
	paises["Argentina"] = "Buenos Aires"

	campeonato := map[string]int{
		"Barcelona":   39,
		"Real Madrid": 38,
		"Chivas":      37,
		"Boca Junior": 30}

	fmt.Println(campeonato)

	campeonato["River Plate"] = 25
	campeonato["Chivas"] = 55
	fmt.Println(campeonato)

}
