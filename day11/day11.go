package day11

import "fmt"

func IncrementString(input []byte) []byte {
	output := []byte{}
	carry := true
	for i := (len(input) - 1); i >= 0; i-- {
		carry = false
		current := input[i]
		current += 1
		if current > 122 {
			carry = true
			current = 97
		}
		output = append([]byte{current}, output...)
		if !carry {
			output = append(input[0:i], output...)
			break
		}
	}
	return output
}

func ContainsThreeLetterSequence(input []byte) bool {
	for k,v := range input {
		if k >= (len(input) - 2) {
			return false
		}
		if v == (input[k+1] - 1) && input[k+1] == (input[k+2] - 1) {
			return true
		}
	}
	return false
}

func ContainsForbiddenLetters(input []byte) bool {
	for _,v := range input {
		if v == 105 || v == 108 || v == 111 {
			return true
		}
	}
	return false
}

func ContainsTwoPairs(input []byte) bool {
	var pairedLetter byte = 0
	pairedPosition := -2
	for k,v := range input {
		if k == (len(input) - 1) {
			return false
		}
		if k == (pairedPosition + 1) {
			continue
		}
		if v == input[k+1] {
			if pairedLetter > 0 {
				return true
			}
			pairedLetter = v
			pairedPosition = k
		}
	}
	return false
}

func GetNextPassword(input []byte) []byte {
	var nextPassword []byte
	for ;; {
		nextPassword = IncrementString(input)
		fmt.Printf("Next possible password: %s\n", nextPassword)
		if !ContainsThreeLetterSequence(nextPassword) {
			continue
		}
		if ContainsForbiddenLetters(nextPassword) {
			continue
		}
		if !ContainsTwoPairs(nextPassword) {
			continue
		}
		return nextPassword
	}
}

func Process(inputString string) {
	nextPassword := GetNextPassword([]byte(inputString))
	fmt.Printf("Next valid password is %s", string(nextPassword[:8]))
}