package day8

import "testing"

func TestEmptyString(t *testing.T) {
	numLiteral,numMemory := parseString(`""`)
	AssertReturnedLength(numLiteral, 2, "literal", t)
	AssertReturnedLength(numMemory, 0, "memory", t)
}

func TestAlphanumericString(t *testing.T) {
	numLiteral,numMemory := parseString(`"abc"`)
	AssertReturnedLength(numLiteral, 5, "literal", t)
	AssertReturnedLength(numMemory, 3, "memory", t)
}

func TestStringWithEscapedQuote(t *testing.T) {
	numLiteral,numMemory := parseString(`"aaa\"aaa"`)
	AssertReturnedLength(numLiteral, 10, "literal", t)
	AssertReturnedLength(numMemory, 7, "memory", t)
}

func TestStringWithEscapedSlash(t *testing.T) {
	numLiteral,numMemory := parseString(`"aaa\\aaa"`)
	AssertReturnedLength(numLiteral, 10, "literal", t)
	AssertReturnedLength(numMemory, 7, "memory", t)
}

func TestStringWithEscapedHex(t *testing.T) {
	numLiteral,numMemory := parseString(`"\x27"`)
	AssertReturnedLength(numLiteral, 6, "literal", t)
	AssertReturnedLength(numMemory, 1, "memory", t)
}

func AssertReturnedLength(value int, expectedValue int, label string, t *testing.T) {
	if value != expectedValue {
		t.Errorf("String %s length should be %d not %d", label, expectedValue, value)
	}
}
