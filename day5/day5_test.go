package day5

import (
	"testing"
)

func TestStringHas3Vowels(t *testing.T) {
	if ! StringHas3Vowels("aeiou") {
		t.Error("aeiou DOES has 3 vowels!")
	}
	if StringHas3Vowels("abcde") {
		t.Error("abcde does NOT has 3 vowels!")
	}
}

func TestStringHasNoRepeatedLetters(t *testing.T) {
	if StringHasRepeatedLetters("aeiou") {
		t.Error("aeiou has NO repeated letters")
	}
}

func TestStringHasRepeatedLetters(t *testing.T) {
	if ! StringHasRepeatedLetters("aabbccdd") {
		t.Error("aabbccdd has NO repeated letters")
	}
}

func TestDoesNotContainABannedString(t *testing.T) {
	if ContainsBannedString("ef") {
		t.Error("ef does NOT contain a banned string")
	}
}

func TestContainsAB(t *testing.T) {
	if ! ContainsBannedString("ab") {
		t.Error("ab is a banned string")
	}
}

func TestContainsCD(t *testing.T) {
	if ! ContainsBannedString("cd") {
		t.Error("cd is a banned string")
	}
}

func TestContainsPQ(t *testing.T) {
	if ! ContainsBannedString("pq") {
		t.Error("pq is a banned string")
	}
}

func TestContainsXY(t *testing.T) {
	if ! ContainsBannedString("xy") {
		t.Error("xy is a banned string")
	}
}

func TestContainsRepeatingPair(t *testing.T) {
	if ! ContainsRepeatingPair("xyxy") {
		t.Error("xyxy DOES contain a repeating pair")
	}
}

func TestContainsOverlappingRepeatingPair(t *testing.T) {
	if ContainsRepeatingPair("aaa") {
		t.Error("aaa does NOT contain a repeating pair")
	}
}

func TestDoesNotContainSandwichLetter(t *testing.T) {
	if ContainsSandwichLetter("abc") {
		t.Error("abc does NOT contain sandwich letter")
	}
}

func TestContainsSandwichLetter(t *testing.T) {
	if ! ContainsSandwichLetter("xyx") {
		t.Error("xyx DOES contain sandwich letter")
	}
}

func TestContainsSandwichLetterWithSameLetterAsFilling(t *testing.T) {
	if ! ContainsSandwichLetter("xxx") {
		t.Error("xxx DOES contain sandwich letter")
	}
}

func TestStringIsNice(t *testing.T) {
	var inputString string
	inputString = "ugknbfddgicrmopn"
	if ! IsNice(inputString) {
		t.Errorf("%s IS a nice string", inputString)
	}

	inputString = "aaa"
	if ! IsNice(inputString) {
		t.Errorf("%s IS a nice string", inputString)
	}

	inputString = "jchzalrnumimnmhp"
	if IsNice(inputString) {
		t.Errorf("%s is NOT a nice string", inputString)
	}

	inputString = "haegwjzuvuyypxyu"
	if IsNice(inputString) {
		t.Errorf("%s is NOT a nice string", inputString)
	}
}

func TestStringIsNicer(t *testing.T) {
	var inputString string
	inputString = "qjhvhtzxzqqjkmpb"
	if ! IsNicer(inputString) {
		t.Errorf("%s IS a nicer string", inputString)
	}

	inputString = "xxyxx"
	if ! IsNicer(inputString) {
		t.Errorf("%s IS a nicer string", inputString)
	}

	inputString = "uurcxstgmygtbstg"
	if IsNicer(inputString) {
		t.Errorf("%s is NOT a nicer string", inputString)
	}

	inputString = "ieodomkazucvgmuy"
	if IsNicer(inputString) {
		t.Errorf("%s is NOT a nicer string", inputString)
	}
}