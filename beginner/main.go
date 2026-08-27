package main

import (
	"fmt"
)

func main() {
	var name string = "Joao"
	fmt.Printf("Say my name, %s\n", name)

	age := 27
	fmt.Printf("%d is not my age...\n", age)

	var city string
	city = "Lisboa"
	fmt.Printf("Cheira a %s\n", city)

	var country, continent string = "PORTUGAL", "Europe"
	fmt.Printf("%s is in %s\n", country, continent)

	var (
		isEmployed bool   = true
		salary     int    = 8001
		position   string = "farmer"
	)

	fmt.Printf("isEmployed : %t \nsalary : %d \nposition : %s\n", isEmployed, salary, position)

	// Zero Values
	var defaultInt int
	var defaultFloat float64
	var defaultString string
	var defaultBool bool

	fmt.Printf("%d ; %f ; '%s' ; %t\n", defaultInt, defaultFloat, defaultString, defaultBool)

	// Consts
	const pi = 3.14

	const (
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
	)

	fmt.Printf("Monday = %d ; Tuesday = %d ; Wednesday = %d\n", Monday, Tuesday, Wednesday)

	const typedAge int = 25
	const untypedAge = 25

	fmt.Println(typedAge == untypedAge)

	const (
		Jan = iota + 1 // 1
		Feb            // 2
		Mar            // 3
		Apr            // 4
	)

	fmt.Printf("jan = %d ; feb = %d ; mar = %d ; apr = %d\n", Jan, Feb, Mar, Apr)

	result := add(2, 2)
	fmt.Printf("The result is: %d\n", result)

	sum, product := calculateSumAndProduct(10, 20)
	fmt.Printf("The sum is %d and the product is %d\n", sum, product)
}

// private function
func add(a int, b int) int {
	return a + b
}

// Public function
func Add(a int, b int) int {
	return a + b
}

func calculateSumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}
