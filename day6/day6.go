package day6

import (
	"regexp"
	"strconv"
	"strings"
	"fmt"
	"time"
)

func TurnOn(row int, col int, lights *[1000][1000]int) {
	lights[row][col] += 1
}

func TurnOff(row int, col int, lights *[1000][1000]int) {
	lights[row][col] -= 1
	if lights[row][col] < 0 {
		lights[row][col] = 0
	}
}

func Toggle(row int, col int, lights *[1000][1000]int) {
	lights[row][col] += 2
}

func Process(instruction string, lights [1000][1000]int, re *regexp.Regexp) [1000][1000]int {
	println("Starting: ", time.Now().Format(time.UnixDate))
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
	println("Parsed params: ", time.Now().Format(time.UnixDate))
	var command string
	if strings.HasPrefix(instruction, "turn on") {
		command = "turn on"
	} else if strings.HasPrefix(instruction, "turn off") {
		command = "turn off"
	} else if strings.HasPrefix(instruction, "toggle") {
		command = "toggle"
	}
	for row := startRow; row <= endRow; row++ {
		println("Starting row: ", time.Now().Format(time.UnixDate))
		for col := startCol; col <= endCol; col++ {
			switch command {
			case "turn on": TurnOn(row, col, &lights)
			case "turn off": TurnOff(row, col, &lights)
			case "toggle": Toggle(row, col, &lights)
			}
		}
		println("Finished row: ", time.Now().Format(time.UnixDate))
	}
	return lights
}

func CountLightsOn(lights [1000][1000]int) int {
	number := 0
	for _,col := range lights {
		for _,light := range col {
			number += light
		}
	}
	return number
}


func Day6(instructions []string) {
	var lights [1000][1000]int
	re := regexp.MustCompile(`[0-9]+`)
	for i,instruction := range instructions {
		fmt.Printf("Processing instruction %d: %s\n", i, instruction)
		lights = Process(instruction, lights, re)
	}
	lightsOn := CountLightsOn(lights)
	fmt.Printf("There are %d lights on", lightsOn)
}