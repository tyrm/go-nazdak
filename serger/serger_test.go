package serger

import "testing"

type testIsInBounds struct {
	x      int
	y      int
	w      int
	h      int
	offset int
	result bool
}

var testIsInBoundss = []testIsInBounds{
	{2,2, 64, 32, 0, true},
	{7,15, 64, 32, 0, true},
	{64,15, 64, 32, 1, true},
	{7,15, 64, 32, 2, false},
	{138,15, 64, 32, 2, true},
}

func TestIsInBounds(t *testing.T) {
	for _, testVal := range testIsInBoundss {

		v := isInBounds(testVal.x, testVal.y, testVal.w, testVal.h, testVal.offset)
		if v != testVal.result {
			t.Error(
				"For", testVal.x, testVal.y, testVal.w, testVal.h, testVal.offset,
				"expected", testVal.result,
				"got", v,
			)
		}
	}
}
