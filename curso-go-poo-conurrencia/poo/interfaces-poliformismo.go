package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Employee
	Person
}

func (fte FullTimeEmployee) getMessage() string {
	return fmt.Sprintf("FullTimeEmployee: %s, %d, %d", fte.name, fte.age, fte.id)
}

type TemporaryEmployee struct {
	Employee
	Person
	taxRate int
}

// Crea una funcion para simular el contrato de implements de otros lenguajes de programación
// Crea y requiere que existan métodos getMessage() string en los structs que implementen la interface PrintInfo
func (te TemporaryEmployee) getMessage() string {
	return fmt.Sprintf("TemporaryEmployee: %s, %d, %d, %d", te.name, te.age, te.id, te.taxRate)
}

// Crea la interface PrintInfo que requiere el método getMessage() string
type PrintInfo interface {
	getMessage() string
}

func getMessage(info PrintInfo) {
	fmt.Println(info.getMessage())
	// return info.getMessage()
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

	tempEmployee := TemporaryEmployee{
		Employee: Employee{id: 3},
		Person:   Person{name: "Alice Doe", age: 28},
		taxRate:  15,
	}

	//Se ejecuta la función getMessage para ambos clases si no se a implementado el contrato en esa clase, se producirá un error en tiempo de compilación
	getMessage(fullEmployee)
	getMessage(tempEmployee)
}
