package main

import (
	"io/ioutil"
	"strings"
	"adventofcode/day12"
)

func main() {
	content, err := ioutil.ReadFile("day12/data.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	day12.Process(lines)
}
