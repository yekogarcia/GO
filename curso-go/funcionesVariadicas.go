package main

import "fmt"

func sumaNumeros(numeros ...int) int {
	fmt.Println("Números recibidos:", numeros)
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	fmt.Println("Total de la suma:", total)
	return total
}

func main() {

	var nombre string = "Yeko"
	var edad int = 25

	fmt.Println("Hola", nombre, "Tengo", edad, "años")

	sumaNumeros(1, 2, 3)
	sumaNumeros(4, 5)
	sumaNumeros(6, 7, 8, 9, 10)

	numeros := []int{11, 12, 13, 14, 15}
	sumaNumeros(numeros...)
}
