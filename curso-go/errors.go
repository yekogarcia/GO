package main

import (
	"errors"
	"fmt"
)

var errorDeCafe = fmt.Errorf("No hay café disponible")
var errorDeEnergia = errors.New("No hay energía disponible")

func hacerCafe(args int) error {
	if args == 2 {
		return errorDeCafe
	} else if args == 4 {
		return errorDeEnergia
	}
	return nil
}

func main() {
	for i := range 5 {
		if err := hacerCafe(i); err != nil {
			if errors.Is(err, errorDeCafe) {
				fmt.Printf("Error al hacer café con args %d: %v\n", i, err)
			} else if errors.Is(err, errorDeEnergia) {
				fmt.Printf("Error al hacer café con args %d: %v\n", i, err)
			} else {
				fmt.Printf("Error desconocido con args %d\n", i, err)
			}
			fmt.Printf("Ya hay cafe en la máquina\n")
		}
	}
}
