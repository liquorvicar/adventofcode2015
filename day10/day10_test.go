package day10

import "testing"

func assertLookAndSayResult(expected string, actual string, t *testing.T) {
	if actual != expected {
		t.Errorf("Expecting %s, got %s", expected, actual)
	}
}

func TestLookAndSay(t *testing.T) {
	assertLookAndSayResult("11", LookAndSay("1"), t)
	assertLookAndSayResult("21", LookAndSay("11"), t)
	assertLookAndSayResult("1211", LookAndSay("21"), t)
	assertLookAndSayResult("111221", LookAndSay("1211"), t)
	assertLookAndSayResult("312211", LookAndSay("111221"), t)
}
