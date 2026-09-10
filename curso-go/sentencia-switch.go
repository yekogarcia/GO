package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Sprintln("i es 1")
	case 2:
		println("i es 2")
	case 3:
		println("i es 3")
	default:
		println("i no es ni 1, ni 2, ni 3")
		println("i es:", i)
		println("Fin del switch")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		println("Es fin de semana")
	default:
		println("Es un día de semana")
	}

	hora := time.Now()
	fmt.Println("Hora actual:", hora)
	switch {
	case hora.Hour() < 12:
		println("Buenos días")
	case hora.Hour() < 18:
		println("Buenas tardes")
	default:
		println("Buenas noches")
	}

}
