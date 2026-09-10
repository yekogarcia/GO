package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int
	x = 8
	y := 7

	fmt.Println(x)
	fmt.Println(y)

	myValue, err := strconv.ParseInt("7", 0, 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed value:", myValue)
	}

	//Maps estructura de clave valor

	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)

	//Slices estructura similar a un arreglo que funciona como una lista
	s := []int{1, 2, 3}

	for index, value := range s {
		fmt.Println("Index:", index, "Value:", value)
	}
	s = append(s, 4)

	fmt.Println(s)
}
