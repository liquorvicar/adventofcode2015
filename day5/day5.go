package day5

import (
	"strings"
	"fmt"
)

func StringHas3Vowels(string string) bool {
	vowels := 0
	vowels += strings.Count(string, "a")
	vowels += strings.Count(string, "e")
	vowels += strings.Count(string, "i")
	vowels += strings.Count(string, "o")
	vowels += strings.Count(string, "u")
	return (vowels >= 3)
}

func StringHasRepeatedLetters(string string) bool {
	var previous rune
	for _, letter := range string {
		if letter == previous {
			return true
		}
		previous = letter
	}
	return false
}

func ContainsBannedString(string string) bool {
	bannedStrings := 0
	bannedStrings += strings.Count(string, "ab")
	bannedStrings += strings.Count(string, "cd")
	bannedStrings += strings.Count(string, "pq")
	bannedStrings += strings.Count(string, "xy")
	return (bannedStrings > 0)
}

func ContainsRepeatingPair(inputString string) bool {
	searchString := inputString
	var previous rune
	firstChar := true
	for _,char := range inputString {
		searchString = strings.TrimPrefix(searchString, string(char))
		if firstChar {
			firstChar = false
		} else {
			pair := string(previous) + string(char)
			if strings.Contains(searchString, pair) {
				return true
			}
		}
		previous = char
	}
	return false
}

func ContainsSandwichLetter(inputString string) bool {
	searchString := inputString
	for _,char := range inputString {
		searchString = strings.TrimPrefix(searchString, string(char))
		if strings.Index(searchString, string(char)) == 1 {
			return true
		} else {
			newSearchString := strings.TrimPrefix(searchString, string(char))
			if strings.Index(newSearchString, string(char)) == 0 {
				return true
			}
		}
	}
	return false
}

func IsNice(string string) bool {
	return StringHas3Vowels(string) && StringHasRepeatedLetters(string) && ! ContainsBannedString(string)
}

func IsNicer(string string) bool {
	return ContainsRepeatingPair(string) && ContainsSandwichLetter(string)
}

func Day5(strings []string) {
	niceStrings := 0
	for _, string := range strings {
		if IsNicer(string) {
			niceStrings += 1
		}
	}
	fmt.Printf("Found %d nice strings\n", niceStrings)
}
