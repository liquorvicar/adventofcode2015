package day2

import "testing"

func TestCalculatePaperForPresent(t *testing.T) {
	volume := calculateVolume(2, 3, 4);
	if volume != 58 {
		t.Errorf("Volume: %d\n", volume);
	}
}

func TestRibbonToWrapPresent(t *testing.T) {
	ribbon := ribbonToWrapPresent(2, 3, 4)
	if ribbon != 10 {
		t.Errorf("Actual ribbon: %d\n", ribbon)
	}
}

func TestRibbonForBow(t *testing.T) {
	ribbon := ribbonForBow(2, 3, 4)
	if ribbon != 24 {
		t.Errorf("Actual ribbon: %d\n", ribbon)
	}
}
