package main

func main() {
	// Ciclo for básico
	for i := 0; i < 5; i++ {
		println("Iteración:", i)
	}

	// Ciclo for con condición
	j := 0
	for j < 5 {
		println("Iteración con condición:", j)
		j++
	}

	// Ciclo for con range
	numeros := []int{1, 2, 3, 4, 5}
	for index, value := range numeros {
		println("Índice:", index, "Valor:", value)
	}

	// Ciclo for infinito (se debe usar con cuidado)
	k := 0
	for {
		if k >= 3 {
			break // Salir del ciclo
		}
		println("Ciclo infinito, iteración:", k)
		k++
	}
}
