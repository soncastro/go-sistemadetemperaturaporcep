package main

import (
	"testing"
)

func TestConvCtoK(t *testing.T) {
	expected := 300.15 // valor de Kelvin esperado para a entrada 27

	result := convCtoK(27)

	if result != expected {
		t.Errorf("convCtoK(27) = %v; want %v", result, expected)
	}
}
