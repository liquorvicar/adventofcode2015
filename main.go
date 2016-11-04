package main

import (
	"io/ioutil"
	"strings"
	"adventofcode2015/day6"
)

func main() {
	content, err := ioutil.ReadFile("day6/data.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	day6.Day6(lines)
}
