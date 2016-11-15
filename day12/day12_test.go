package day12

import "testing"

func assertCount(actual int, expected int, t *testing.T) {
	if actual != expected {
		t.Errorf("Sum should be %d, it is actually %d", expected, actual)
	}
}

func TestCountNumbers(t *testing.T) {
	assertCount(CountNumbers(`[1,2,3]`), 6, t)
	assertCount(CountNumbers(`{"a":2,"b":4}`), 6, t)
	assertCount(CountNumbers(`[[[3]]]`), 3, t)
	assertCount(CountNumbers(`{"a":{"b":4},"c":-1}`), 3, t)
	assertCount(CountNumbers(`{"a":[-1,1]}`), 0, t)
	assertCount(CountNumbers(`[-1,{"a":1}]`), 0, t)
	assertCount(CountNumbers(`[]`), 0, t)
	assertCount(CountNumbers(`{}`), 0, t)
}
