package day6

import (
	"testing"
	"regexp"
)

func TestTurnOnLight(t *testing.T) {
	lights := [1000][1000]int{{0}}
	lights = TurnOn(0, 0, lights)
	if lights[0][0] != 1 {
		t.Error("Light should be turned on")
	}
}

func TestTurnOffLight(t *testing.T) {
	lights := [1000][1000]int{{1}}
	lights = TurnOff(0, 0, lights)
	if lights[0][0] != 0 {
		t.Error("Light should be turned off")
	}
}

func TestToggleLightOff(t *testing.T) {
	lights := [1000][1000]int{{1}}
	lights = Toggle(0, 0, lights)
	if lights[0][0] != 3 {
		t.Error("Light should be toggled off")
	}
}

func TestToggleLightOn(t *testing.T) {
	lights := [1000][1000]int{{0}}
	lights = Toggle(0, 0, lights)
	if lights[0][0] != 2 {
		t.Error("Light should be toggled on")
	}
}

func TestProcessTurnOnInstruction(t *testing.T) {
	lights := [1000][1000]int{{0}}
	re := regexp.MustCompile(`[0-9]+`)
	lights = Process("turn on 0,0 through 2,2", lights, re)
	if lights[0][0] != 1 {
		t.Error("Light should be on")
	}
	if lights[1][0] != 1 {
		t.Error("Light should be on")
	}
	if lights[2][0] != 1 {
		t.Error("Light should be on")
	}
	if lights[0][1] != 1 {
		t.Error("Light should be on")
	}
	if lights[1][1] != 1 {
		t.Error("Light should be on")
	}
	if lights[1][2] != 1 {
		t.Error("Light should be on")
	}
	if lights[2][0] != 1 {
		t.Error("Light should be on")
	}
	if lights[2][1] != 1 {
		t.Error("Light should be on")
	}
	if lights[2][2] != 1 {
		t.Error("Light should be on")
	}
}

func TestProcessTurnOffInstruction(t *testing.T) {
	lights := [1000][1000]int{{0}}
	lights[0][0] = 1
	re := regexp.MustCompile(`[0-9]+`)
	lights = Process("turn off 0,0 through 0,0", lights, re)
	if lights[0][0] != 0 {
		t.Error("Light should be off")
	}
}

func TestProcessToggleInstruction(t *testing.T) {
	lights := [1000][1000]int{{0}}
	lights[0][0] = 1
	lights = Process("toggle 0,0 through 0,0", lights, regexp.MustCompile(`[0-9]+`))
	if lights[0][0] != 3 {
		t.Error("Light should be off")
	}
}

func TestCountLightsOn(t *testing.T) {
	lights := [1000][1000]int{{0}}
	lights = Process("turn on 0,0 through 2,2", lights, regexp.MustCompile(`[0-9]+`))
	lightsOn := CountLightsOn(lights)
	if lightsOn != 9 {
		t.Errorf("There should be 9 lights on, not %d", lightsOn)
	}
}
