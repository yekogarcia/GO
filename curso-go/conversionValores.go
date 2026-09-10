package main

import "fmt"

// Conversion of values in Go

func main() {
	fmt.Println("Conversion of values in Go")

	var entero int = 10
	entero2 := 20
	fmt.Println(entero + entero2)

	var numeroEntero = 30
	var numeroDoble = 10.6

	resultado := numeroEntero + int(numeroDoble)
	fmt.Println("Resultado de la suma:", resultado)

	var nombre string = "yeko"
	apellido := "garcia"
	nombreCompleto := nombre + " " + apellido
	fmt.Println("Nombre completo:", nombreCompleto)

}
