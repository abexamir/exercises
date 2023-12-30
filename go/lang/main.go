package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	var reader = bufio.NewScanner(os.Stdin)

	type Person struct {
		name   string
		grades []int
	}

	var people []Person

	for i := 1; i <= number; i++ {
		var person Person
		reader.Scan()
		person.name = reader.Text()
		reader.Scan()
		input := reader.Text()
		grades := strings.Fields(input)
		for _, numStr := range grades {
			num, _ := strconv.Atoi(numStr)
			person.grades = append(person.grades, num)
		}
		people = append(people, person)
	}
	for _, person := range people {
		average := calculateAverage(person.grades)
		switch {
		case average >= 80:
			fmt.Printf("%s %s\n", person.name, "Excelent")
		case average >= 60:
			fmt.Printf("%s %s\n", person.name, "Very Good")
		case average >= 40:
			fmt.Printf("%s %s\n", person.name, "Good")
		default:
			fmt.Printf("%s %s\n", person.name, "Fair")
		}
	}
}
