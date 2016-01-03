package strategy

import "testing"

func TestMultiplicationStrategy(t *testing.T) {
	mult := Operation{Multiplication{}}

	if res := mult.Operate(3, 5); res != 15 {
		t.Errorf("Multiplication.Operate(3, 5) expected 15 got %q", res)
	}
}

func TestAdditionStrategy(t *testing.T) {
	add := Operation{Addition{}}
	if res := add.Operate(3, 5); res != 8 {
		t.Errorf("Addition.Operate(3, 5) expected 8 got %q", res)
	}
}
