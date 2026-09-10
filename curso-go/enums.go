package main

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "Idle",
	StateConnected: "Connected",
	StateError:     "Error",
	StateRetrying:  "Retrying",
}

func (estado ServerState) String() string {
	return stateName[estado]
}

func main() {
	redServidor := verificacionDeRed(StateIdle)
	fmt.Println("Estado del servidor:", redServidor)

	segRevision := verificacionDeRed(redServidor)
	fmt.Println("Estado del servidor después de la revisión:", segRevision)

	tercRevision := verificacionDeRed(segRevision)
	fmt.Println("Estado del servidor después de la tercera revisión:", tercRevision)

	cuartaRevision := verificacionDeRed(tercRevision)
	fmt.Println("Estado del servidor después de la cuarta revisión:", cuartaRevision)

	// Intento de pasar un estado desconocido
	// Esto causará un panic
	// verificacionDeRed(ServerState(999))
}

func verificacionDeRed(servidor ServerState) ServerState {
	switch servidor {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("Estado desconocido del servidor: %v", servidor))
	}
}
