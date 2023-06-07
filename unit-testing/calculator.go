package main

type Calculator struct {
}

func (c Calculator) Multiply(a int, b int) int {
	return a * b
}

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Divide(a, b int) int {
	return a / b
}
