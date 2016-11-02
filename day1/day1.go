package main

import (
	"io/ioutil"
	"fmt"
)

func main() {

	rawContents, err := ioutil.ReadFile("data/day1.txt")
	if err != nil {
		panic(err)
	}
	instructions := string(rawContents[:])
	floor := 0
	for n, instruction := range instructions {
		switch instruction {
		case '(': floor += 1;
		case ')': floor -= 1;
		}
		if floor < 0 {
			fmt.Printf("Stopping at instruction: %d\n", (n +1));
			break;
		}
	}
	fmt.Printf("Final floor: %d", floor);
}
