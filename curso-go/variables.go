package main

import "fmt"

func main() {
	// Declare and initialize variables

	var name string = "John"
	var age int = 30
	var height float64 = 5.9
	var entero1, entero2 int = 10, 20
	var isboolean bool = true
	var enteroSimple int
	enteroSimple = 15

	fruta := "apple" // Short variable declaration

	// Print the variables
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Entero1 Entero2:", entero1, entero2)
	fmt.Println("isBoolean:", isboolean)
	fmt.Println("enteroSimple:", enteroSimple)
	fmt.Println("Fruit:", fruta)

	// Declare and initialize a constant
	const pi float64 = 3.14159
	fmt.Println("Value of Pi:", pi)
}
