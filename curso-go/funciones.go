package main

import "fmt"

func suma(a int, b int) int {
	return a + b
}

func sumaTresNumeros(a, b, c int) int {
	return a + b + c
}

func main() {

	var a, b, c int
	fmt.Print("Ingrese el primer número: ")
	fmt.Scan(&a)
	fmt.Print("Ingrese el segundo número: ")
	fmt.Scan(&b)
	fmt.Print("Ingrese el tercer número: ")
	fmt.Scan(&c)

	resultado := sumaTresNumeros(a, b, c)
	fmt.Println("La suma de los tres números es:", resultado)

	total := suma(a, b)
	fmt.Println("La suma de los dos primeros números es:", total)

}
