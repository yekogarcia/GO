package main

import (
	"fmt"
	"os"
)

func main() {
	enVar := os.Getenv("HOME")
	if enVar == "" {
		fmt.Println("HOME environment variable is not set.")
	} else {
		fmt.Println("Home:", enVar)
	}
}
