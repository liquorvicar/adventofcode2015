package main

import (
	"io/ioutil"
	"strings"
	"adventofcode2015/day7"
)

func main() {
	content, err := ioutil.ReadFile("day7/data.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	day7.Day7(lines)
}
