package main

import (
	"fmt"
	"slices"
)

func main() {
	var arregloCadenas []string

	arregloCadenas = make([]string, 3)

	arregloCadenas[0] = "Hola"
	arregloCadenas[1] = "Mundo"
	arregloCadenas[2] = "Go"

	arregloCadenas = append(arregloCadenas, "Programación")
	fmt.Println("Arreglo de cadenas:", arregloCadenas)

	segundoArreglo := []string{"Go", "es", "genial"}
	fmt.Println("Segundo arreglo:", segundoArreglo)

	tercerArreglo := []string{"Go", "es", "el mejor lenguaje de programación"}
	fmt.Println("Tercer arreglo:", tercerArreglo)

	if slices.Equal(segundoArreglo, tercerArreglo) {
		fmt.Println("Los arreglos son iguales")
	} else {
		fmt.Println("Los arreglos son diferentes")
	}

	// fmt.Println("Arreglo de cadenas:", arregloCadenas, "condición:", arregloCadenas == nil, "longitud:", len(arregloCadenas) == 0)
}
