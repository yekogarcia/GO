package main

import "fmt"

type Persona struct {
	Nombre string
	Edad   int
}

type Empleado struct {
	Persona
	Cargo string
}

func nuevaPersona(nombre string) *Persona {
	nuevaIndividuo := Persona{Nombre: nombre}
	nuevaIndividuo.Edad = 42
	return &nuevaIndividuo
}

func main() {

	fmt.Println(Persona{"Ender", 25})

	fmt.Println(Persona{Nombre: "Santiago", Edad: 30})

	fmt.Println(Empleado{Persona: Persona{Nombre: "Santiago", Edad: 30}, Cargo: "Desarrollador"})

	fmt.Println(Empleado{Persona: Persona{Nombre: "Santiago"}, Cargo: "Desarrollador"})

	fmt.Println(nuevaPersona("Roberto"))

	personita := Persona{Nombre: "Miranda", Edad: 30}
	fmt.Println(personita.Nombre)

	edadPersonita := &personita
	fmt.Println(edadPersonita.Edad)

	edadPersonita.Edad = 35
	fmt.Println(edadPersonita.Edad)
	fmt.Println(personita)
}
