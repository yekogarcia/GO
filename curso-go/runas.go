package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	const saludo = "สวัสดีชาวโลก"

	fmt.Println("El saludo en tailandes es:", saludo)
	fmt.Println("Len:", len(saludo))

	for i := 0; i < len(saludo); i++ {
		fmt.Printf("%x ", saludo[i])
	}

	fmt.Println("Rune count:", utf8.RuneCountInString(saludo))

	for idx, runeValue := range saludo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, idx)
	}
}
