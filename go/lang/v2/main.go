package main

import (
	"fmt"
)

type Person struct {
	name   string
	grades []int
}

func calculateAverage(grades []int) int {
	sum := 0
	for _, grade := range grades {
		sum += grade
	}
	return sum / len(grades)
}

func main() {
	var number int
	fmt.Scanf("%d", &number)

	people := make([]Person, number)

	for i := 0; i < number; i++ {
		var person Person
		fmt.Scanln(&person.name)
		var grade int
		var scanning bool = true
		for scanning {
			_, err := fmt.Scanf("%d", &grade)
			if err != nil {
				scanning = false
			} else {
				person.grades = append(person.grades, grade)
			}
		}
		people[i] = person
	}

	for _, person := range people {
		average := calculateAverage(person.grades)
		switch {
		case average >= 80:
			fmt.Printf("%s %s\n", person.name, "Excellent")
		case average >= 60:
			fmt.Printf("%s %s\n", person.name, "Very Good")
		case average >= 40:
			fmt.Printf("%s %s\n", person.name, "Good")
		default:
			fmt.Printf("%s %s\n", person.name, "Fair")
		}
	}
}
