package main

import "fmt"

func modificarNumeros(arg *[5]int) {
	(*arg)[0] = 42
}

func main() {

	numeros := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Números antes de modificar:", numeros)

	modificarNumeros(&numeros)
	fmt.Println("Números después de modificar:", numeros)
}
