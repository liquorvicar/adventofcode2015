package main

import (
	"adventofcode2015/day5"
	"io/ioutil"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("day5/data.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	day5.Day5(lines)
}
