package day8

import (
	"strings"
	"unicode/utf8"
	"regexp"
	"fmt"
)

func parseString(inputString string) (int,int) {
	hex := regexp.MustCompile(`\\x[0-9a-f]{2}`)
	parsedString := strings.Trim(inputString, `"`)
	parsedString = strings.Replace(parsedString, `\"`, `Q`, -1)
	parsedString = strings.Replace(parsedString, `\\`, `S`, -1)
	parsedString = hex.ReplaceAllString(parsedString, `X`)
	return utf8.RuneCountInString(inputString),utf8.RuneCountInString(parsedString)
}

func Process(lines []string) {
	answer := 0
	for _,inputString := range lines {
		numLiteral,numMemory := parseString(inputString)
		fmt.Printf("String %s has %d literals and %d chars\n", inputString, numLiteral, numMemory)
		answer += (numLiteral - numMemory)
	}
	fmt.Printf("Answer is %d\n", answer)
 }