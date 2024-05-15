package main

import (
	"math"
	"testing"
)

func TestConvKELtoC(t *testing.T) {
	expected := 26.85 // o valor em Celsius esperado para a entrada 300

	result := convKELtoC(300)

	if math.Abs(result-expected) > 0.01 {
		t.Errorf("convKELtoC(300) = %v; want %v", result, expected)
	}
}

func TestConvKELtoF(t *testing.T) {
	expected := 80.33 // o valor em Fahrenheit esperado para a entrada 300

	result := convKELtoF(300)

	if math.Abs(result-expected) > 0.01 {
		t.Errorf("convKELtoF(300) = %v; want %v", result, expected)
	}
}
