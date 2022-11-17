package Calculator

import (
	"errors"
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	testsAdd := []struct {
		num1     float64
		num2     float64
		response float64
	}{
		{1, 2, 3},
		{1.0, 3, 4},
		{-1, 1, 0},
		{-2.2, -2.2, -4.4},
	}

	for i,tc:= range testsAdd {
		actualOp := add(tc.num1, tc.num2)
		if reflect.DeepEqual(actualOp,tc.response) {
			t.Log("Correct output")
		} else {
			t.Errorf("Test no. %v: Wanted (%f) but got", i, tc.response)
		}
	}
}
func TestSubtract(t *testing.T) {
	testsSubtract := []struct {
		num1     float64
		num2     float64
		response float64
	}{
		{2, 1, 0},
		{3.0, 1, 2},
		{-1.0, 1.0, -2},
		{-2, -2, 0},
	}

	for i,tc := range testsSubtract {
		actualOp := subtract(tc.num1, tc.num2)
		if reflect.DeepEqual(actualOp, testsSubtract[i].response) {
			t.Log("Correct output")
		} else {
			t.Errorf("Test no. %v: Wanted (%f) but got(%f)", i, tc.response, actualOp)
		}
	}
}

func TestMultiply(t *testing.T) {
	testsMultiply := []struct {
		num1     float64
		num2     float64
		response float64
	}{
		{2, 1, 2},
		{3.0, 1, 3},
		{-1.0, 1.0, -1},
		{-2, -2, 4},
		{1.2, 1.2, 1.44},
	}

	for i,tc := range testsMultiply {
		actualOp := multiply(tc.num1, tc.num2)
		if reflect.DeepEqual(actualOp, testsMultiply[i].response) {
			t.Log("Correct output")
		} else {
			t.Errorf("input at %v index has expected value as %f and actual output as %f", i, tc.response, actualOp)
		}
	}
}

func TestDivision(t *testing.T) {
	testsDivision := []struct {
		num1     float64
		num2     float64
		response float64
		err      error
	}{
		{2, 1, 2, nil},
		{3.0, 1, 3, nil},
		{-1.0, 1.0, -1, nil},
		{-2, -2, 1, nil},
		{1.2, 1.2, 1.44, nil},
		{1, 0, 0, errors.New("divide by zero not possible")},
	}

	for i,tc := range testsDivision {
		actualOp, err := division(tc.num1, tc.num2)
		if err == nil && err == testsDivision[i].err {
			if reflect.DeepEqual(actualOp, tc.response) {
				t.Log("Correct output")
			} else {
				t.Errorf("input at %v index has expected value as %f and actual output as %f", i, tc.response, actualOp)
			}
		} else {
			t.Errorf("input at %v index has expected error as %e and actual error as %e", i, tc.err, err)
		}
	}
}

func TestModulo(t *testing.T) {
	testsModulo := []struct {
		num1     int
		num2     int
		response int
		err      error
	}{
		{2, 1, 0, nil},
		{3.0, 1, 0, nil},
		{-1.0, 0, 0, errors.New("modulo by zero not possible")},
		{-2, -2, 0, nil},
	}

	for i,tc:= range testsModulo {
		actualOp, err := modulo(tc.num1, tc.num2)
		if err == nil && err == tc.err {
			if reflect.DeepEqual(actualOp, tc.response) {
				t.Log("Correct output")
			} else {
				t.Errorf("input at %v index has expected value as %d and actual output as %d", i, tc.response, actualOp)
			}
		} else {
			t.Errorf("input at %v index has expected error as %d and actual error as %d", i, testsModulo[i].err, err)
		}
	}
}
func TestCalculator(t *testing.T) {
	tests := []struct {
		num1     float64
		num2     float64
		op       string
		response float64
		err      error
	}{
		{2, 1, "+", 3, nil},
		{3.0, 1, "-", 2, nil},
		{-1.0, 0, "*", 0, nil},
		{-2, -2, "/", 1, nil},
		{4, 2, "%", 0, nil},
		{3, 1, "&", 0, errors.New("invalid operator")},
		{3, 1, "!", 0, errors.New("invalid operator")},
		{3, 1, "@", 0, errors.New("invalid operator")},
		{3, 1, "#", 0, errors.New("invalid operator")},
	}

	for i := range tests {
		actualOutput, err := calculator(tests[i].num1, tests[i].num2, tests[i].op)
		if err == nil && err == tests[i].err {
			if reflect.DeepEqual(actualOutput, tests[i].response) {
				t.Log("Correct output")
			} else {
				t.Errorf("input at %v index has expected value as %f and actual output as %f", i, tests[i].response, actualOutput)
			}
		} else {
			t.Errorf("input at %v index has expected error as %d and actual error as %d", i, tests[i].err, err)
		}

	}
}
