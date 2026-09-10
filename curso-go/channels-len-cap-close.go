package main

import (
	"fmt"
)

func main() {

	c := make(chan string, 2)
	c <- "Mensaje1"
	c <- "Mensaje2"

	fmt.Println(len(c), cap(c))

	close(c)

	for message := range c {
		fmt.Println(message)
	}

	//Select

	email1 := make(chan string)
	email2 := make(chan string)
	go messages("Email 1", email1)
	go messages("Email 2", email2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-email1:
			fmt.Println("Email 1 Recibido:", msg1)
		case msg2 := <-email2:
			fmt.Println("Email 2 Recibido:", msg2)
		}
	}

}

func messages(text string, c chan string) {
	c <- text
}
