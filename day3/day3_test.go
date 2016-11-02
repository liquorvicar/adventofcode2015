package day3

import "testing"

func TestOneHouseDeliveredTo(t *testing.T) {
	houses := calculateHousesDeliveredTo(">")
	if houses != 2 {
		t.Errorf("Delivered to %d houses \n", houses)
	}
}

func TestFourHousesDeliveredTo(t *testing.T) {
	houses := calculateHousesDeliveredTo("^>v<")
	if houses != 4 {
		t.Errorf("Delivered to %d houses \n", houses)
	}
}

func TestTwoHousesDeliveredToWithRobo(t *testing.T) {
	houses := calculateHousesDeliveredToWithRobo("^v")
	if houses != 3 {
		t.Errorf("Delivered to %d houses \n", houses)
	}
}

func TestElevenHousesDeliveredToWithRobo(t *testing.T) {
	houses := calculateHousesDeliveredToWithRobo("^v^v^v^v^v")
	if houses != 11 {
		t.Errorf("Delivered to %d houses \n", houses)
	}
}
