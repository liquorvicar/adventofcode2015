package day9

import (
	"regexp"
	"strings"
	"strconv"
	"fmt"
)

var citiesRE = regexp.MustCompile(`([a-zA-Z]+) to ([a-zA-Z]+)`)
func parseCities(inputStrings []string) map[string]bool {
	allCities := make(map[string]bool)
	for _,inputString := range inputStrings {
		cities := citiesRE.FindAllStringSubmatch(inputString, -1)
		_,ok := allCities[cities[0][1]]
		if ! ok {
			allCities[cities[0][1]] = true
		}
		_,ok = allCities[cities[0][2]]
		if ! ok {
			allCities[cities[0][2]] = true
		}
	}
	return allCities
}

var distancesRE = regexp.MustCompile(`([a-zA-Z]+) to ([a-zA-Z]+) = ([0-9]+)`)
func allDistances(inputStrings []string) map[string]int {
	allDistances := make(map[string]int)
	for _,inputString := range inputStrings {
		cities := distancesRE.FindAllStringSubmatch(inputString, -1)
		distance,err := strconv.Atoi(cities[0][3])
		if err == nil {
			allDistances[cities[0][1] + "," + cities[0][2]] = distance
			allDistances[cities[0][2] + "," + cities[0][1]] = distance
		}
	}
	return allDistances
}

func allRoutes(destinations map[string]bool) map[string]int {
	routes := make(map[string]int)
	parseRoutes(destinations, "", routes)
	return routes
}

func parseRoutes(destinations map[string]bool, currentRoute string, routes map[string]int) {
	if len(destinations) == 0 {
		currentRoute = strings.TrimLeft(currentRoute, "-")
		routes[currentRoute] = 0
		return
	}
	for destination := range destinations {
		newDestinations := make(map[string]bool)
		for k,v := range destinations {
			if k != destination {
				newDestinations[k] = v
			}
		}
		parseRoutes(newDestinations, currentRoute + "-" + string(destination), routes)
	}
	return
}

func shortestDistance(allRoutes map[string]int, allDistances map[string]int) int {
	shortest := -1
	for route := range allRoutes {
		distance := 0
		destinations := strings.Split(route, "-")
		for i,destination := range destinations {
			if i == 0 {
				continue
			}
			journey := destinations[i-1] + "," + destination
			thisDistance,_ := allDistances[journey]
			distance += thisDistance
		}
		if shortest < 0 || distance < shortest {
			shortest = distance
		}
	}
	return shortest
}

func longestDistance(allRoutes map[string]int, allDistances map[string]int) int {
	longest := 0
	for route := range allRoutes {
		distance := 0
		destinations := strings.Split(route, "-")
		for i,destination := range destinations {
			if i == 0 {
				continue
			}
			journey := destinations[i-1] + "," + destination
			thisDistance,_ := allDistances[journey]
			distance += thisDistance
		}
		if distance > longest {
			longest = distance
		}
	}
	return longest
}

func Process(inputStrings []string) {
	allRoutes := allRoutes(parseCities(inputStrings))
	allDistances := allDistances(inputStrings)
	shortest := longestDistance(allRoutes, allDistances)
	fmt.Printf("Longest distance is %d\n", shortest)

}