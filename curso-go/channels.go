package main

func main() {
	mensajes := make(chan string)

	go func() {
		mensajes <- "Hola desde la gorutina"
	}()

	mensaje := <-mensajes
	println(mensaje)

	go func() {
		mensajes <- "Hola desde la gorutina 2"
	}()

	mensaje2 := <-mensajes
	println(mensaje2)
}
