package main

import "fmt"

func valoresMultipes(nombre1, nombre2, apellido string) (map[string]string, map[string]string) {
	nombres := make(map[string]string)
	nombres[nombre1] = nombre1
	nombres[nombre2] = nombre2

	fullName := make(map[string]string)
	fullName[nombre1] = nombre1
	fullName[nombre2] = nombre2
	fullName[apellido] = apellido

	return nombres, fullName
}

func main() {
	nombre1 := "Ender"
	nombre2 := "Yecid"
	apellido := "Garcia"

	valoresMultipes1, valoresMultipes2 := valoresMultipes(nombre1, nombre2, apellido)
	fmt.Println("Valores múltiples 1:", valoresMultipes1)
	fmt.Println("Valores múltiples 2:", valoresMultipes2)

	if valoresMultipes1["nombre"] == valoresMultipes2["nombre"] && valoresMultipes1["apellido"] == valoresMultipes2["apellido"] {
		fmt.Println("Los mapas son iguales")
	} else {
		fmt.Println("Los mapas son diferentes")
	}

	_, nombre := valoresMultipes(nombre1, nombre2, apellido)
	fmt.Println("Valor de 'nombre':", nombre)

}
