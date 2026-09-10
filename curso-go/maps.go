package main

import (
	"fmt"
	"maps"
)

func main() {
	// Declaración de un slice de enteros
	mapa := make(map[string]int)

	// Asignación de valores al mapa
	mapa["uno"] = 1
	mapa["dos"] = 2
	mapa["tres"] = 3

	fmt.Println("Mapa:", mapa)

	version1 := mapa["uno"]
	fmt.Println("Valor de 'uno':", version1)

	// Excluir un elemento del mapa
	_, datoExiste := mapa["tres"]
	if datoExiste {
		fmt.Println("El dato 'tres' existe en el mapa")
	} else {
		fmt.Println("El dato 'tres' no existe en el mapa")
	}

	// Eliminar un elemento del mapa
	delete(mapa, "dos")
	fmt.Println("Mapa después de eliminar 'dos':", mapa)

	//Limpiar el mapa
	clear(mapa)
	fmt.Println("Mapa después de limpiar:", mapa)

	newMaps1 := make(map[string]int)
	newMaps1["cuatro"] = 4
	newMaps1["cinco"] = 5
	fmt.Println("Nuevo mapa:", newMaps1)

	newMaps2 := map[string]int{
		"seis":  6,
		"siete": 7,
	}
	fmt.Println("Otro nuevo mapa:", newMaps2)

	if maps.Equal(newMaps1, newMaps2) {
		fmt.Println("Los mapas son iguales")
	} else {
		fmt.Println("Los mapas son diferentes")
	}

}
