package main

import "fmt"

type Employe struct {
	id   int
	name string
	age  int
}

// Aca se crea una funcion para simular como la de las contrcuores en otros lenguajes.
// La función se puede llamar normalmente como cualquier otra función.
func constructor(id int, name string, age int) *Employe {
	return &Employe{
		id:   id,
		name: name,
		age:  age,
	}
}

func main() {
	// Aca se crea una instancia usando la funcion constructor y asignando los valores.
	e := constructor(1, "John", 30)
	fmt.Printf("%v", e)

	// Se puede acceder también de esta forma.
	e1 := Employe{
		id:   1,
		name: "edwar",
		age:  22,
	}
	fmt.Printf("%v", e1)

	e2 := new(Employe)
	fmt.Printf("%v", e2)

	e2.age = 35
	e2.name = "Alice"
	e2.id = 2

	fmt.Printf("%v", e2)
}
