package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	texto := "Hola, mundo!"
	hash1 := sha256.Sum256([]byte(texto))
	fmt.Printf("SHA-256: %x\n", hash1)

	hash2 := sha256.New()
	hash2.Write([]byte(texto))
	arregloHash2 := hash2.Sum(nil)

	fmt.Println("SHA-256: %x\n", arregloHash2)
	fmt.Printf("SHA-256: %x\n", arregloHash2)
}
