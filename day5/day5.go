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

func IsNice(string string) bool {
	return StringHas3Vowels(string) && StringHasRepeatedLetters(string) && ! ContainsBannedString(string)
}

func Day5(strings []string) {
	niceStrings := 0
	for _, string := range strings {
		if IsNice(string) {
			niceStrings += 1
		}
	}
	fmt.Printf("Found %d nice strings\n", niceStrings)
}
