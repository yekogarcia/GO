package main

import "fmt"

// class Person.
type Person struct {
	name string
	age  int
}

// class Employee.
type Employee struct {
	id int
}

// class FullTimeEmployee. Simula la herencia llamando los otros structs Employee y Person. que serian las simulaciones de clases en otro lenguaje.
// Lo que hace realmente es composición: es combinar los campos de Employee y Person en un solo struct.
type FullTimeEmployee struct {
	Employee
	Person
}

func main() {
	fullEmployee := FullTimeEmployee{
		Employee: Employee{id: 1},
		Person:   Person{name: "John Doe", age: 30},
	}
	fmt.Println(fullEmployee)

	fullEmployee2 := FullTimeEmployee{}
	fmt.Printf("%+v\n", fullEmployee2)

	fullEmployee2.age = 25
	fullEmployee2.name = "Jane Doe"
	fullEmployee2.id = 2

	fmt.Println(fullEmployee2)

}
