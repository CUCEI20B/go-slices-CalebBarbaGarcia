package main

import "fmt"

func main() {
	var s []int
	var tamano uint
	var aux int
	var i int
	var i1 uint
	var resultado int

	
	fmt.Scanln(&tamano)
	i1 = 0
	
	for i1 < tamano {
		fmt.Scanln(&aux)
		s = append(s,aux)
		i1 = i1 + 1
	}

	i = 0

	for i < len(s) {
		resultado = resultado + s[i]
		i = i + 1
	}

	println(resultado)
}
