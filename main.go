package main

import "fmt"

func main() {
	greetin := "Hello there, my first code of GO :D \n"
	fmt.Print(greetin)

	// slide
	numbers := []int{3, 4}

	numbers = append(numbers, 6, 7)
	fmt.Print(numbers)

	numbers = numbers[3:] // take a range of numbers
	fmt.Print(numbers)

	// for en go, this is the only one that exists
	// Example 1
	for i := range numbers {
		fmt.Print(numbers[i])
	}
	// Example 2
	for i, number := range numbers {
		fmt.Print("\n Index: ", i, " Value: ", number)
	}
	// Example 3
	for _, number := range numbers {
		fmt.Print("\n Value: ", number)
	}

	fruits := []string{
		"Manzana",
		"Banano",
		"Sandia",
		"Melon",
	}

	for _, fruit := range fruits {
		if fruit != "Manzana" {
			continue
		}

		fmt.Print("\n Fruit: ", fruit)
	}

}
