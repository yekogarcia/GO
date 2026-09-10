package main

import (
	"fmt"
	"strconv"
	"time"
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

	//Se crea un Channel de tipo int para comunicarse con la gorutine
	c := make(chan int)
	go doSomething(c)
	<-c

	g := 25
	fmt.Println(g)
	//Crea la referencia a la variable g o el apuntamiento a la dirección de memoria de g
	h := &g
	fmt.Println(h)
	//con * se accede al valor al que apunta la referencia
	fmt.Println(*h)
}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Doing something")
	c <- 1
}
