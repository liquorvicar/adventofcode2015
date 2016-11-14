package day10

import (
	"strconv"
	"fmt"
)

func LookAndSay(inputString string) string {
	outputString := ""
	currentValue := string(inputString[0])
	count := 0
	for _,char := range inputString {
		if currentValue == string(char) {
			count += 1
			continue
		}
		outputString += strconv.Itoa(count) + string(currentValue)
		currentValue = string(char)
		count = 1
	}
	outputString += strconv.Itoa(count) + string(currentValue)
	return outputString
}

func RepeatLookAndSay(inputString string, count int) int {
	for i := 1; i <= count; i++ {
		inputString = LookAndSay(inputString)
	}
	return len(inputString)
}

func Process(inputString string) {
	fmt.Printf("Length of final string is %d\n", RepeatLookAndSay(inputString, 40))
}