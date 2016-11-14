package day11

import "testing"

func assertIncrementedString(expected string, actual []byte, t *testing.T) {
	if string(actual[:]) != expected {
		t.Errorf("Expected %s but got %s", expected, string(actual[:2]))
	}
}

func TestIncrementString(t *testing.T) {
	assertIncrementedString("ba", IncrementString([]byte("az")), t)
	assertIncrementedString("ya", IncrementString([]byte("xz")), t)
	assertIncrementedString("abd", IncrementString([]byte("abc")), t)
	assertIncrementedString("abcdefgi", IncrementString([]byte("abcdefgh")), t)
}

func TestContainsThreeLetterSequence(t *testing.T) {
	if !ContainsThreeLetterSequence([]byte("abc")) {
		t.Error("abc does contain three letter sequence")
	}
	if !ContainsThreeLetterSequence([]byte("cde")) {
		t.Error("cde does contain three letter sequence")
	}
	if ContainsThreeLetterSequence([]byte("abd")) {
		t.Error("abd does not contain three letter sequence")
	}
	if !ContainsThreeLetterSequence([]byte("xxcdexx")) {
		t.Error("xxcdexx does contain three letter sequence")
	}
}

func TestContainsForbiddenLetters(t *testing.T) {
	if !ContainsForbiddenLetters([]byte("ghi")) {
		t.Error("ghi contains a forbidden letter")
	}
	if !ContainsForbiddenLetters([]byte("jkl")) {
		t.Error("jkl contains a forbidden letter")
	}
	if !ContainsForbiddenLetters([]byte("mno")) {
		t.Error("mno contains a forbidden letter")
	}
	if ContainsForbiddenLetters([]byte("abc")) {
		t.Error("abc does not contain a forbidden letter")
	}
}

func TestContainsTwoPairs(t *testing.T) {
	if !ContainsTwoPairs([]byte("aabb")) {
		t.Error("aabb does contain two pairs")
	}
	if !ContainsTwoPairs([]byte("aaxyzbb")) {
		t.Error("aaxyzbb does contain two pairs")
	}
	if !ContainsTwoPairs([]byte("fgaaxyzbblm")) {
		t.Error("fgaaxyzbblm does contain two pairs")
	}
	if ContainsTwoPairs([]byte("aaa")) {
		t.Error("aaa does not contain two pairs")
	}
	if ContainsTwoPairs([]byte("fgaaxyzblm")) {
		t.Error("fgaaxyzblm does not contain two pairs")
	}
}

func TestGivenExamples(t *testing.T) {
	if !ContainsThreeLetterSequence([]byte("hijklmmn")) {
		t.Error("hijklmmn does contain three letter sequence")
	}
	if !ContainsForbiddenLetters([]byte("hijklmmn")) {
		t.Error("hijklmmn does contain forbidden chars")
	}
	if ContainsThreeLetterSequence([]byte("abbceffg")) {
		t.Error("abbceffg does not contain three letter sequence")
	}
	if !ContainsTwoPairs([]byte("abbceffg")) {
		t.Error("abbceffg does contain two pairs")
	}
}

func TestGetNextPassword(t *testing.T) {
	if string(GetNextPassword([]byte("abcdefgh"))[:8]) != "abcdffaa" {
		t.Error("Next password after abcdefgh should be abcdffaa")
	}
}
