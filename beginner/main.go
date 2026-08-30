package main

import (
	"fmt"
)

func main() {
	age := 30

	if age >= 18 {
		fmt.Println("You can drive!!!!!!")
	} else if age >= 13 {
		fmt.Println("You suck")
	} else {
		fmt.Println("You need more Danoninho")
	}

	day := "Friday"

	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("Midweek")
	case "Friday":
		fmt.Println("ITS OVER!!!!")
		fallthrough
	default:
		fmt.Println("Rest for now")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Look: ", i)
	}

	counter := 0

	for counter < 3 {
		fmt.Println("Look again: ", counter)
		counter++
	}

	iter := 0
	for {
		if iter > 3 {
			break
		}
		iter++
	}

	numbers := [5]int{1, 2, 3, 4, 5}

	fmt.Printf("The array %v\n", numbers)

	// allNumbers := numbers[:]
	// firstThree := numbers[0:3]

	fruits := []string{"apple", "banana", "strawberry"}
	fmt.Printf("these are my fruits %v\n", fruits)

	fruits = append(fruits, "kiwi")
	fmt.Printf("these are my fruits with kiwi %v\n", fruits)

	fruits = append(fruits, "mango", "pineapple")
	fmt.Printf("these are my fruits with more fruits%v\n", fruits)

	moreFruits := []string{"watermelon", "melon"}
	fruits = append(fruits, moreFruits...)
	fmt.Printf("these are my fruits with more fruits%v\n", fruits)

	for index, value := range numbers {
		fmt.Printf("[%d] %d\n", index, value)
	}

	capitalCities := map[string]string{
		"Portugal": "Lisbon",
		"USA":      "Washington D.C.",
		"UK":       "London",
	}

	capital, exists := capitalCities["Germany"]
	if exists {
		fmt.Println("this is the capital", capital)
	} else {
		fmt.Println("Does not exist")
	}

	delete(capitalCities, "USA")
	fmt.Printf("this is new deleted map: %v\n", capitalCities)
}

// private function
func add(a int, b int) int {
	return a + b
}

func calculateSumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}
