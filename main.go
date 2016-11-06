package main

import (
	"io/ioutil"
	"strings"
	"adventofcode2015/day9"
)

func main() {
	content, err := ioutil.ReadFile("day9/data.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	day9.Process(lines)
}
