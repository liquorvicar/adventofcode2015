package day9

import (
	"testing"
)

var testInput = []string{"London to Dublin = 464", "London to Belfast = 518", "Dublin to Belfast = 141"}

func TestAllDestinations(t *testing.T) {
	allDestinations := parseCities(testInput)
	if len(allDestinations) != 3 {
		t.Error("Parsed list of cities is incorrect")
	}
	if _,ok := allDestinations["Belfast"]; ! ok {
		t.Error("Parsed list of cities is incorrect")
	}
	if _,ok := allDestinations["Dublin"]; ! ok {
		t.Error("Parsed list of cities is incorrect")
	}
	if _,ok := allDestinations["London"]; ! ok {
		t.Error("Parsed list of cities is incorrect")
	}
}

func TestAllRoutes(t *testing.T) {
	allDestinations := make(map[string]bool)
	allDestinations["Belfast"] = true
	allDestinations["Dublin"] = true
	allDestinations["London"] = true
	allRoutes := allRoutes(allDestinations)
	if len(allRoutes) != 6 {
		t.Error("Parsed list of routes is incorrect")
	}
	if _,ok := allRoutes["Belfast-Dublin-London"]; ! ok {
		t.Error("Parsed list of routes is incorrect")
	}
	if _,ok := allRoutes["London-Dublin-Belfast"]; ! ok {
		t.Error("Parsed list of routes is incorrect")
	}
	if _,ok := allRoutes["Belfast-London-Dublin"]; ! ok {
		t.Error("Parsed list of routes is incorrect")
	}
	if _,ok := allRoutes["Dublin-London-Belfast"]; ! ok {
		t.Error("Parsed list of routes is incorrect")
	}
	if _,ok := allRoutes["London-Belfast-Dublin"]; ! ok {
		t.Error("Parsed list of routes is incorrect")
	}
	if _,ok := allRoutes["Dublin-Belfast-London"]; ! ok {
		t.Error("Parsed list of routes is incorrect")
	}
}

func TestAllDistances(t *testing.T) {
	allDistances := allDistances(testInput)
	if len(allDistances) != 6 {
		t.Error("Parsed list of distances is incorrect")
	}
	if distance,ok := allDistances["Dublin,London"]; ! ok || distance != 464 {
		t.Error("Parsed list of distances is incorrect")
	}
	if distance,ok := allDistances["London,Dublin"]; ! ok || distance != 464 {
		t.Error("Parsed list of distances is incorrect")
	}
	if distance,ok := allDistances["Belfast,London"]; ! ok || distance != 518 {
		t.Error("Parsed list of distances is incorrect")
	}
	if distance,ok := allDistances["London,Belfast"]; ! ok || distance != 518 {
		t.Error("Parsed list of distances is incorrect")
	}
	if distance,ok := allDistances["Dublin,Belfast"]; ! ok || distance != 141 {
		t.Error("Parsed list of distances is incorrect")
	}
	if distance,ok := allDistances["Belfast,Dublin"]; ! ok || distance != 141 {
		t.Error("Parsed list of distances is incorrect")
	}
}

func TestFindShortestDistance(t *testing.T) {
	allRoutes := allRoutes(parseCities(testInput))
	allDistances := allDistances(testInput)
	shortest := shortestDistance(allRoutes, allDistances)
	if shortest != 605 {
		t.Errorf("Distance of %d is not 605!", shortest)
	}
}