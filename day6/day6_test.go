package day6

import "testing"

func TestTurnOnLight(t *testing.T) {
	lights := [1000][1000]bool{{false}}
	lights = TurnOn(0, 0, lights)
	if lights[0][0] == false {
		t.Error("Light should be turned on")
	}
}

func TestTurnOffLight(t *testing.T) {
	lights := [1000][1000]bool{{true}}
	lights = TurnOff(0, 0, lights)
	if lights[0][0] == true {
		t.Error("Light should be turned off")
	}
}

func TestToggleLightOff(t *testing.T) {
	lights := [1000][1000]bool{{true}}
	lights = Toggle(0, 0, lights)
	if lights[0][0] == true {
		t.Error("Light should be toggled off")
	}
}

func TestToggleLightOn(t *testing.T) {
	lights := [1000][1000]bool{{false}}
	lights = Toggle(0, 0, lights)
	if lights[0][0] == false {
		t.Error("Light should be toggled on")
	}
}

func TestProcessTurnOnInstruction(t *testing.T) {
	lights := [1000][1000]bool{{false}}
	lights = Process("turn on 0,0 through 2,2", lights)
	if lights[0][0] == false {
		t.Error("Light should be on")
	}
	if lights[1][0] == false {
		t.Error("Light should be on")
	}
	if lights[2][0] == false {
		t.Error("Light should be on")
	}
	if lights[0][1] == false {
		t.Error("Light should be on")
	}
	if lights[1][1] == false {
		t.Error("Light should be on")
	}
	if lights[1][2] == false {
		t.Error("Light should be on")
	}
	if lights[2][0] == false {
		t.Error("Light should be on")
	}
	if lights[2][1] == false {
		t.Error("Light should be on")
	}
	if lights[2][2] == false {
		t.Error("Light should be on")
	}
}

func TestProcessTurnOffInstruction(t *testing.T) {
	lights := [1000][1000]bool{{false}}
	lights[0][0] = true
	lights = Process("turn off 0,0 through 0,0", lights)
	if lights[0][0] == true {
		t.Error("Light should be off")
	}
}

func TestProcessToggleInstruction(t *testing.T) {
	lights := [1000][1000]bool{{false}}
	lights[0][0] = true
	lights = Process("toggle 0,0 through 0,0", lights)
	if lights[0][0] == true {
		t.Error("Light should be off")
	}
}

func TestCountLightsOn(t *testing.T) {
	lights := [1000][1000]bool{{false}}
	lights = Process("turn on 0,0 through 2,2", lights)
	lightsOn := CountLightsOn(lights)
	if lightsOn != 9 {
		t.Errorf("There should be 9 lights on, not %d", lightsOn)
	}
}
