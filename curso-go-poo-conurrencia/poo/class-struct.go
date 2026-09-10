package main

import "fmt"

// Los structs son los mismos que las clases en otros lenguajes, pero sin métodos especiales como constructores.
type Employe struct {
	id   int
	name string
	age  int
}

// Set
func (e *Employe) SetId(id int) {
	e.id = id
}

func (e *Employe) SetName(name string) {
	e.name = name
}

func (e *Employe) SetAge(age int) {
	e.age = age
}

// Get
func (e *Employe) GetId() int {
	return e.id
}

func (e *Employe) GetName() string {
	return e.name
}

func (e *Employe) GetAge() int {
	return e.age
}

func main() {
	e := Employe{id: 1, name: "John", age: 30}
	fmt.Printf("%v", e)

	e.SetId(2)
	e.SetName("Jane")
	e.SetAge(25)

	fmt.Println("ID:", e.GetId())
	fmt.Println("Name:", e.GetName())
	fmt.Println("Age:", e.GetAge())

}
