package main

import "fmt"

func main() {
	// Declaración de un arreglo de enteros con 5 elementos
	var numeros [5]int

	// Asignación de valores a los elementos del arreglo
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros[3] = 40
	numeros[4] = 50

	// Acceso a los elementos del arreglo
	fmt.Println("Primer elemento:", numeros[0])
	fmt.Println("Segundo elemento:", numeros[1])
	fmt.Println("Tercer elemento:", numeros[2])
	fmt.Println("Cuarto elemento:", numeros[3])
	fmt.Println("Quinto elemento:", numeros[4])

	// Longitud del arreglo
	fmt.Println("Longitud del arreglo:", len(numeros))
	fmt.Println("Arreglo:", numeros)

	lista := [3]string{"Manzana", "Banana", "Cereza"}
	fmt.Println("Lista de frutas:", lista)

	arrayUnlimitedo := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Arreglo con tamaño inferido:", arrayUnlimitedo)
}
