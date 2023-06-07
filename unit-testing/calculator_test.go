package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Expected %d, but get %d", expected, result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	expected := 2
	if result != expected {
		t.Errorf("Expected %d, "+
			"but get %d", expected, result)
	}
}

func TestMultiply(t *testing.T) {
	calculator := Calculator{}
	result := calculator.Multiply(4, 2)
	expected := 8
	if result != expected {
		t.Errorf("Expected %d, but get %d", expected, result)
	}
}

func TestDivide(t *testing.T) {
	result := Divide(6, 3)
	expected := 2
	if result != expected {
		t.Errorf("Expected %d, but get %d", expected, result)
	}
}

func TestAdd2(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	assert.Equal(t, expected, result, "Hasil penjumlahan tidak sesuai")
}
