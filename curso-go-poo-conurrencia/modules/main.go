package main

import (
	"fmt"

	"github.com/donvito/hellomod"
	hellomodv2 "github.com/donvito/hellomod/v2"
)

func main() {
	fmt.Println("Hello, Modules!")
	hellomod.SayHello()
	hellomodv2.SayHello("Platzi")
}
