package main

import (
	"fmt"
	"time"
)

func main() {
	canal1 := make(chan string)
	canal2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second) // Simular trabajo
		canal1 <- "Mensaje desde canal 1"
	}()

	go func() {
		// time.Sleep(2 * time.Second) // Simular trabajo
		canal2 <- "Mensaje desde canal 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case mensaje1 := <-canal1:
			fmt.Println("Recibido:", mensaje1)
		case mensaje2 := <-canal2:
			fmt.Println("Recibido:", mensaje2)
		}
	}

}
