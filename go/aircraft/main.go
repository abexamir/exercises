package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func number_of_arithmetic_slices(record []int) int {
	if len(record) < 3 {
		return 0
	}
	count := 0
	counting := false
	main_diff := record[1] - record[0]
	i := 1
	for i < len(record)-1 {
		second_diff := record[i+1] - record[i]
		if second_diff == main_diff {
			if !counting {
				counting = true
				count += 1
			} else {
				count += 1
			}
		} else {
			counting = false
			main_diff = second_diff
		}
		i++
	}
	return count
}

func main() {
	var number int
	// fmt.Scanf("%d", &number)
	number = 2
	scanner := bufio.NewScanner(os.Stdin)
	var names []string
	var records [][]int
	for i := 0; i < number; i++ {
		scanner.Scan()
		tokens := strings.Fields(scanner.Text())
		name := tokens[0]
		record := make([]int, len(tokens)-1)
		for j := 1; j < len(tokens); j++ {
			record[j-1], _ = strconv.Atoi(tokens[j])
		}
		names = append(names, name)
		records = append(records, record)
	}
	for i := 0; i < number; i++ {
		fmt.Printf("%s %d\n", names[i], number_of_arithmetic_slices(records[i]))
	}
}
