package day3

import "fmt"

type deliveryRecord struct {
	x int
	y int
}

func calculateHousesDeliveredTo(instructions string) int {
	var houses = make(map[int]map[int]bool)
	var santa = deliveryRecord {
		x: 0,
		y: 0,
	}
	houses[0] = make(map[int]bool)
	houses[0][0] = true
	for _, instruction := range instructions {
		deliverTo(&santa, houses, instruction)
	}
	return countHouses(houses)
}

func calculateHousesDeliveredToWithRobo(instructions string) int {
	var houses = make(map[int]map[int]bool)
	var santa = deliveryRecord {
		x: 0,
		y: 0,
	}
	var robo = deliveryRecord {
		x: 0,
		y: 0,
	}
	useSanta := true
	houses[0] = make(map[int]bool)
	houses[0][0] = true
	for _, instruction := range instructions {
		if useSanta {
			deliverTo(&santa, houses, instruction)
		} else {
			deliverTo(&robo, houses, instruction)
		}
		useSanta = !useSanta
	}
	return countHouses(houses)
}

func countHouses(houses map[int]map[int]bool) int {
	var countDelivered = 0
	for _, row := range houses {
		for _, delivered := range row {
			if delivered {
				countDelivered += 1
			}
		}
	}
	return countDelivered
}

func deliverTo(santa *deliveryRecord, houses map[int]map[int]bool, instruction rune) {
	switch instruction {
	case '^': santa.y += 1
	case 'v': santa.y -= 1
	case '>': santa.x += 1
	case '<': santa.x -= 1
	}
	_,ok := houses[santa.x]
	if ! ok {
		houses[santa.x] = make(map[int]bool)
	}
	houses[santa.x][santa.y] = true

}

func Day3(lines []string) {
	deliveryCount := 0
	for _,instructions := range lines {
		deliveryCount += calculateHousesDeliveredToWithRobo(instructions)
	}
	fmt.Printf("Delivered to %d houses\n", deliveryCount)
}