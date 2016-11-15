package day12

import (
	"regexp"
	"strconv"
	"fmt"
)

func CountNumbers(input string) int {
	re := regexp.MustCompile(`(-?[0-9]+)`)
	matches := re.FindAllString(input, -1)
	sum := 0
	for _,value := range matches {
		number,err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("Cannot convert %s to int\n", value)
		} else {
			sum += number
		}
	}
	return sum
}

func Process(input []string) {
	for _,data := range input {
		println(CountNumbers(data))
	}
}