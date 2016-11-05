package day7

import (
	"regexp"
	"strconv"
	"fmt"
)

type Wire struct {
	value uint16
	processed bool
}

func (wire Wire) Value() uint16 {
	return wire.value
}

type Gate struct {
	input1, input2 Wire
}

type Command interface {
	Value() uint16
}

type AndGate Gate

func (gate AndGate) Value() uint16 {
	return gate.input1.Value() & gate.input2.Value()
}

type OrGate Gate

func (gate OrGate) Value() uint16 {
	return gate.input1.Value() | gate.input2.Value()
}

type LeftShiftGate Gate

func (gate LeftShiftGate) Value() uint16 {
	return gate.input1.Value() << gate.input2.Value()
}

type RightShiftGate Gate

func (gate RightShiftGate) Value() uint16 {
	return gate.input1.Value() >> gate.input2.Value()
}

type NotGate Gate

func (gate NotGate) Value() uint16 {
	return ^gate.input1.Value()
}

func Process(instruction string, circuit map[string]Wire) bool {
	if UnmatchedOperation(instruction) {
		panic("Unmatched operation: " + instruction)
	}
	return ProcessValue(instruction, circuit) ||
	ProcessAndOr(instruction, circuit) ||
	ProcessAndOrWithValue(instruction, circuit) ||
	ProcessShift(instruction, circuit) ||
	ProcessNot(instruction, circuit) ||
	ProcessWireToWire(instruction, circuit)
}

func UnmatchedOperation(instruction string) bool {
	return valueRE.FindStringSubmatch(instruction) == nil &&
	andOrRE.FindStringSubmatch(instruction) == nil &&
	andOrWithValueRE.FindStringSubmatch(instruction) == nil &&
	shiftRE.FindStringSubmatch(instruction) == nil &&
	notRE.FindStringSubmatch(instruction) == nil &&
	wireToWireRE.FindStringSubmatch(instruction) == nil
}

var valueRE = regexp.MustCompile(`^([0-9]+) -> ([a-z]+)$`)
func ProcessValue(instruction string, circuit map[string]Wire) bool {
	parts := valueRE.FindStringSubmatch(instruction)
	if parts == nil {
		return false
	}
	inputValue, err := strconv.Atoi(parts[1])
	if err != nil {
		return false
	}
	label := parts[2]
	wire := Wire{value: uint16(inputValue), processed: true}
	circuit[label] = wire
	return true
}

var wireToWireRE = regexp.MustCompile(`^([a-z]+) -> ([a-z]+)$`)
func ProcessWireToWire(instruction string, circuit map[string]Wire) bool {
	parts := wireToWireRE.FindStringSubmatch(instruction)
	if parts == nil {
		return false
	}
	inputValue, ok := circuit[parts[1]]
	if ! ok || ! inputValue.processed {
		return false
	}
	label := parts[2]
	wire := Wire{value: inputValue.Value(), processed: true}
	circuit[label] = wire
	return true
}

var andOrRE = regexp.MustCompile(`([a-z]+) ([A-Z]+) ([a-z]+) -> ([a-z]+)`)
func ProcessAndOr(instruction string, circuit map[string]Wire) bool {
	parts := andOrRE.FindStringSubmatch(instruction)
	if parts == nil {
		return false
	}
	input1,ok := circuit[parts[1]]
	if ! ok || ! input1.processed {
		return false
	}
	input2,ok := circuit[parts[3]]
	if ! ok || ! input2.processed {
		return false
	}
	label := parts[4]
	command := parts[2]
	var gate Command
	switch command {
	case "AND": gate = AndGate{input1, input2}
	case "OR": gate = OrGate{input1, input2}
	}

	wire := Wire{value: gate.Value(), processed: true}
	circuit[label] = wire
	return true
}

var andOrWithValueRE = regexp.MustCompile(`([0-9]+) ([A-Z]+) ([a-z]+) -> ([a-z]+)`)
func ProcessAndOrWithValue(instruction string, circuit map[string]Wire) bool {
	parts := andOrWithValueRE.FindStringSubmatch(instruction)
	if parts == nil {
		return false
	}
	input1,err := strconv.Atoi(parts[1])
	if err != nil {
		return false
	}
	input2,ok := circuit[parts[3]]
	if ! ok || ! input2.processed {
		return false
	}
	label := parts[4]
	command := parts[2]
	var gate Command
	switch command {
	case "AND": gate = AndGate{Wire{uint16(input1), true}, input2}
	case "OR": gate = OrGate{Wire{uint16(input1), true}, input2}
	}

	wire := Wire{value: gate.Value(), processed: true}
	circuit[label] = wire
	return true
}

var shiftRE = regexp.MustCompile(`([a-z]+) ([A-Z]+) ([0-9]+) -> ([a-z]+)`)
func ProcessShift(instruction string, circuit map[string]Wire) bool {
	parts := shiftRE.FindStringSubmatch(instruction)
	if parts == nil {
		return false
	}
	input1,ok := circuit[parts[1]]
	if ! ok || ! input1.processed {
		return false
	}
	input2,err := strconv.Atoi(parts[3])
	if err != nil {
		return false
	}
	label := parts[4]
	command := parts[2]
	var gate Command
	switch command {
	case "LSHIFT": gate = LeftShiftGate{input1, Wire{uint16(input2), true}}
	case "RSHIFT": gate = RightShiftGate{input1, Wire{uint16(input2), true}}
	}

	wire := Wire{value: gate.Value(), processed: true}
	circuit[label] = wire
	return true
}

var notRE = regexp.MustCompile(`NOT ([a-z]+) -> ([a-z]+)`)
func ProcessNot(instruction string, circuit map[string]Wire) bool {
	parts := notRE.FindStringSubmatch(instruction)
	if parts == nil {
		return false
	}
	input1,ok := circuit[parts[1]]
	if ! ok || ! input1.processed {
		return false
	}
	label := parts[2]
	gate := NotGate{input1, Wire{uint16(0), true}}

	wire := Wire{value: gate.Value(), processed: true}
	circuit[label] = wire
	return true
}

func Day7(instructions []string) {
	circuit := map[string]Wire{}
	processed := make(map[string]bool)
	finished := false
	iterations := 0
	for ; ! finished; {
		finished = true
		numProcessed := 0
		for _,instruction := range instructions {
			hasBeenProcessed,exists := processed[instruction]
			if exists && hasBeenProcessed {
				continue
			}
			thisWasProcessed := Process(instruction, circuit)
			if thisWasProcessed {
				println("Processed: ", instruction)
				processed[instruction] = true
				numProcessed += 1
			}
			finished = finished && thisWasProcessed
		}
		iterations += 1
		fmt.Printf("Iteration %d, %d processed\n", iterations, numProcessed)
		if numProcessed == 0 {
			finished = true
		}
	}
	println("++++ Finished ++++")
	answer,ok := circuit["a"]
	if !ok {
		panic("a not found")
	//	for _,instruction := range instructions {
	//		hasBeenProcessed, exists := processed[instruction]
	//		if ! exists || ! hasBeenProcessed {
	//			println(instruction)
	//		}
	//	}
	}
	println("++++ Finished ++++")
	fmt.Printf("Wire a has final value %d", answer.Value())
}