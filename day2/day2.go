package day2

import (
	"strings"
	"fmt"
	"strconv"
	"sort"
)

func calculateVolume(h, w, l int) int {
	sides := [3]int { (h * w), (h * l), (l * w)}
	volume := 0
	smallest := -1
	for _, side := range sides {
		volume += (2 * side)
		if (smallest < 0) || (smallest > side) {
			smallest = side
		}
	}
	return (volume + smallest)
}

func ribbonToWrapPresent(h, w, l int) int {
	dims := []int {h, w, l}
	sort.Ints(dims)
	return (dims[0] * 2) + (dims[1] * 2)
}

func ribbonForBow(h, w, l int) int {
	return (h * w * l)
}

func Day2(lines []string) {
	totalRibbon := 0
	for _, line := range lines {
		values := strings.Split(line, "x")
		h, err := strconv.Atoi(values[0])
		if err != nil {
			panic(err)
		}
		w, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}
		l, err := strconv.Atoi(values[2])
		if err != nil {
			panic(err)
		}
		totalRibbon += ribbonToWrapPresent(h, w, l)
		totalRibbon += ribbonForBow(h, w, l)
	}
	fmt.Printf("Total ribbon: %d\n", totalRibbon)
}
