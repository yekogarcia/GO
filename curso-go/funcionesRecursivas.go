package main

import "fmt"

// Función recursiva para calcular el factorial de un número
// 7! = 7 * 6 * 5 * 4 * 3 * 2 * 1
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
func main() {
	var numero int
	fmt.Print("Ingrese un número para calcular su factorial: ")
	fmt.Scan(&numero)
	fmt.Println("El factorial de", numero, "es", factorial(numero))

	// Función recursiva para calcular el Fibonacci de un número
	// Fibonacci: 0, 1, 1, 2, 3, 5, 8, 13, 21, ...
	var finobacci func(numeroBase int) int
	finobacci = func(n int) int {
		if n <= 1 {
			return n
		}
		return finobacci(n-1) + finobacci(n-2)
	}

	fmt.Print("Ingrese un número para calcular su Fibonacci: ")
	fmt.Scan(&numero)
	fmt.Println("El Fibonacci de", numero, "es", finobacci(numero))
}
