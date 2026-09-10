package main

import (
	"fmt"
	"time"
)

func function(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		// time.Sleep(100 * time.Millisecond)
	}
}

func main() {

	function("modo directo")

	go function("modo gorutine")

	go func(mensaje string) {
		fmt.Println(mensaje)
	}("mensaje desde una función anónima")

	time.Sleep(1 * time.Second) // Esperar a que las gorutinas terminen antes de salir del programa
	fmt.Println("fin del programa")

}
