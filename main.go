package main

import (
	"io/ioutil"
	"strings"
	"adventofcode2015/day8"
)

func main() {
	content, err := ioutil.ReadFile("day8/data.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	day8.Process(lines)
}
