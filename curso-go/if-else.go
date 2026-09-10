package main

import "fmt"

func main() {

	nombre := "yeko"
	edad := 28

	if nombre == "yeko" {
		fmt.Println("El nombre es Yeko")
	} else {
		println("Hola desconocido")
	}

	if edad > 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}

	if 8%2 == 0 {
		fmt.Println("El número es par")
	} else {
		fmt.Println("El número es impar")
	}

	if numero := 99; numero < 0 {
		fmt.Println("El número es negativo")
	} else if numero > 0 {
		fmt.Println("El número es positivo")
	} else {
		fmt.Println("El número es cero")
	}

}
