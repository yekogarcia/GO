package main

import "fmt"

func main() {
	fmt.Println(sum(3, 4, 5, 1))
	printNames("Alice", "Bob", "Charlie")
	fmt.Println(getValue(3))
	fmt.Println(getValueVar(3))

}

// variadic function
func sum(values ...int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

// variadic function
func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

// returns
func getValue(x int) (int, int, int) {
	return x, x * 2, x * 3
}

func getValueVar(x int) (double int, triple int, quadruple int) {
	double = x * 2
	triple = x * 3
	quadruple = x * 4
	return
}
