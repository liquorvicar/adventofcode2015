package day6

import (
	"regexp"
	"strconv"
	"strings"
	"fmt"
)

func TurnOn(row int, col int, lights [1000][1000]bool) [1000][1000]bool {
	lights[row][col] = true
	return lights
}

func TurnOff(row int, col int, lights [1000][1000]bool) [1000][1000]bool {
	lights[row][col] = false
	return lights
}

func Toggle(row int, col int, lights [1000][1000]bool) [1000][1000]bool {
	lights[row][col] = ! lights[row][col]
	return lights
}

func Process(instruction string, lights [1000][1000]bool, re *regexp.Regexp) [1000][1000]bool {
	bounds := re.FindAllString(instruction, -1)
	startRow, err := strconv.Atoi(bounds[0])
	if err != nil {
		panic("Invalid start row")
	}
	startCol, err := strconv.Atoi(bounds[1])
	if err != nil {
		panic("Invalid end row")
	}
	endRow, err := strconv.Atoi(bounds[2])
	if err != nil {
		panic("Invalid start col")
	}
	endCol, err := strconv.Atoi(bounds[3])
	if err != nil {
		panic("Invalid end col")
	}
	var command string
	if strings.HasPrefix(instruction, "turn on") {
		command = "turn on"
	} else if strings.HasPrefix(instruction, "turn off") {
		command = "turn off"
	} else if strings.HasPrefix(instruction, "toggle") {
		command = "toggle"
	}
	for row := startRow; row <= endRow; row++ {
		for col := startCol; col <= endCol; col++ {
			switch command {
			case "turn on": lights = TurnOn(row, col, lights)
			case "turn off": lights = TurnOff(row, col, lights)
			case "toggle": lights = Toggle(row, col, lights)
			}
		}
	}
	return lights
}

func CountLightsOn(lights [1000][1000]bool) int {
	number := 0
	for _,col := range lights {
		for _,light := range col {
			if light {
				number += 1
			}
		}
	}
	return number
}

func Day6(instructions []string) {
	var lights [1000][1000]bool
	re := regexp.MustCompile(`[0-9]+`)
	for i,instruction := range instructions {
		fmt.Printf("Processing instruction %d: %s\n", i, instruction)
		lights = Process(instruction, lights, re)
	}
	lightsOn := CountLightsOn(lights)
	fmt.Printf("There are %d lights on", lightsOn)
}