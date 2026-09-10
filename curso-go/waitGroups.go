package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Iniciando gorutina...")
	wg.Add(1)

	go say("Hola desde la gorutina", &wg)

	wg.Wait()

	go func(text string) {
		fmt.Println(text)
	}("Hola desde la gorutina 2")
}
