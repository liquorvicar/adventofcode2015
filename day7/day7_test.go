package day7

import "testing"

func TestValueInputToWire(t *testing.T) {
	wire := Wire{value: 123}
	if wire.Value() != 123 {
		t.Errorf("Value is %d but should be 123", wire.Value())
	}
}

func TestAndGate(t *testing.T) {
	x := Wire{value: 5}
	y := Wire{value: 6}
	and := AndGate{x, y}
	if and.Value() != 4 {
		t.Errorf("Value is %d but should be 4", and.Value())
	}
}

func TestOrGate(t *testing.T) {
	x := Wire{value: 2}
	y := Wire{value: 4}
	or := OrGate{x, y}
	if or.Value() != 6 {
		t.Errorf("Value is %d but should be 6", or.Value())
	}
}

func TestLeftShiftGate(t *testing.T) {
	x := Wire{value: 1}
	y := Wire{value: 2}
	leftShift := LeftShiftGate{x, y}
	if leftShift.Value() != 4 {
		t.Errorf("Value is %d but should be 6", leftShift.Value())
	}
}

func TestRightShiftGate(t *testing.T) {
	x := Wire{value: 4}
	y := Wire{value: 2}
	rightShift := RightShiftGate{x, y}
	if rightShift.Value() != 1 {
		t.Errorf("Value is %d but should be 1", rightShift.Value())
	}
}

func TestNotGate(t *testing.T) {
	x := Wire{value: 65535}
	y := Wire{value: 0}
	not := NotGate{x, y}
	if not.Value() != 0 {
		t.Errorf("Value is %d but should be 0", not.Value())
	}
}

func AssertWire(circuit map[string]Wire, label string, value uint16, t *testing.T) {
	if circuit[label].Value() != value {
		t.Errorf("Value of %s should be %d, actually: %d", label, value, circuit[label].Value())
	}
}

func TestCircuit(t *testing.T) {
	circuit := map[string]Wire{}
	Process("123 -> x", circuit)
	Process("456 -> y", circuit)
	Process("x AND y -> d", circuit)
	Process("x OR y -> e", circuit)
	Process("x LSHIFT 2 -> f", circuit)
	Process("y RSHIFT 2 -> g", circuit)
	Process("NOT x -> h", circuit)
	Process("NOT y -> i", circuit)
	AssertWire(circuit, "x", 123, t)
	AssertWire(circuit, "y", 456, t)
	AssertWire(circuit, "d", 72, t)
	AssertWire(circuit, "e", 507, t)
	AssertWire(circuit, "f", 492, t)
	AssertWire(circuit, "g", 114, t)
	AssertWire(circuit, "h", 65412, t)
	AssertWire(circuit, "i", 65079, t)
}