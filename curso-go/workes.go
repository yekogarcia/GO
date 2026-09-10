package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d: procesando trabajo %d\n", id, j)
		time.Sleep(time.Second) // Simular trabajo
		fmt.Println("Worker", id, "Tarea iniciada", j)
		results <- j * 2
	}
}

func main() {
	const numWorkers = 5
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 10; a++ {
		result := <-results
		fmt.Printf("Resultado recibido: %d\n", result)
	}
}
