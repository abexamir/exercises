package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var number int
	fmt.Scanf("%d", &number)

	scanner := bufio.NewScanner(os.Stdin)
	countries := make(map[string]string)
	areaCodes := make([]string, 0)

	for i := 0; i < number; i++ {
		scanner.Scan()
		countryCode := scanner.Text() // Iran +98
		tokens := strings.Split(countryCode, " ")
		countries[tokens[1]] = tokens[0]
	}
	var phoneNumbersCount int
	fmt.Scanf("%d", &phoneNumbersCount)
	for i := 0; i < phoneNumbersCount; i++ {
		scanner.Scan()
		phoneNumber := scanner.Text()
		areaCode := phoneNumber[:3]
		areaCodes = append(areaCodes, areaCode)
	}

	for _, v := range areaCodes {
		country, ok := countries[v]
		if ok {
			fmt.Printf("%s\n", country)
		} else {
			fmt.Printf("Invalid Number\n")
		}
	}
}
