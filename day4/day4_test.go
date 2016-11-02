package day4

import "testing"

func TestBasicMD5(t *testing.T) {
	md5 := findMatchingMD5("abcdef")
	if md5 != 609043 {
		t.Errorf("Incorrect md5 %v\n", md5)
	}
}

func TestAnotherMD5(t *testing.T) {
	md5 := findMatchingMD5("pqrstuv")
	if md5 != 1048970 {
		t.Errorf("Incorrect md5 %v\n", md5)
	}
}
