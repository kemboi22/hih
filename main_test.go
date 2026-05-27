package main

import "testing"

func TestAdd(t *testing.T) {
	if got := Add(1, 2); got != 3 {
		t.Errorf("Add(1, 2) = %d; want 3", got)
	}
}

func TestSubtract(t *testing.T) {
	if got := Subtract(5, 3); got != 2 {
		t.Errorf("Subtract(5, 3) = %d; want 2", got)
	}
}

func TestMultiply(t *testing.T) {
	if got := Multiply(4, 3); got != 12 {
		t.Errorf("Multiply(4, 3) = %d; want 12", got)
	}
}

func TestDivide(t *testing.T) {
	if got := Divide(10, 2); got != 5 {
		t.Errorf("Divide(10, 2) = %d; want 5", got)
	}
}

func TestDivideByZero(t *testing.T) {
	if got := Divide(10, 0); got != 0 {
		t.Errorf("Divide(10, 0) = %d; want 0", got)
	}
}
