package main

import (
	"encoding/json"
	"fmt"
)

type Persona struct {
	Nombre string `json:"nombre"`
	Edad   int    `json:"edad"`
}

func main() {
	persona := Persona{Nombre: "Santiago", Edad: 30}

	jsonData, err := json.Marshal(persona)
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}

	fmt.Println("JSON:", string(jsonData))

	var persona2 Persona
	err = json.Unmarshal(jsonData, &persona2)
	if err != nil {
		fmt.Println("Error al convertir de JSON a objeto:", err)
		return
	}

	fmt.Printf("Persona desde Object: %+v\n", persona2)
}
