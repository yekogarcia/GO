package main

import (
	"fmt"
	"time"
)

func main() {

	contador := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second) // Simular trabajo
		contador <- "Resultado 1"
	}()

	select {
	case mensaje := <-contador:
		fmt.Println("Recibido:", mensaje)
	case <-time.After(2 * time.Second):
		fmt.Println("Tiempo de espera agotado")
	}
}
