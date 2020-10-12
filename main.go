package main

import "fmt"

func main() {
	var s []int
	var tamano int
	var aux int
	var i int
	var resultado int

	
	fmt.Scan(&tamano)
	i = 0
	for i < tamano {
		fmt.Scan(&aux)
		s = append(s,aux)
		i = i + 1
	}

	i = 0

	for i < len(s) {
		resultado = resultado + s[i]
		i = i + 1
	}

	println(resultado)
}
