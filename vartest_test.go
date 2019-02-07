package vartest

import (
	"testing"
)

func TestWithVariable(t *testing.T) {

	a, b := WithVariable()
	if !(a == 8 && b == 10) {
		panic("err")
	}
}

func TestWithoutVariable(t *testing.T) {

	a, b := WithoutVariable()
	if !(a == 8 && b == 10) {
		panic("err")
	}
}

func TestWithArray(t *testing.T) {

	a, b := WithArray()
	if !(a == 55 && b == -55) {
		panic("err")
	}
}

func TestWithoutArray(t *testing.T) {

	a, b := WithoutArray()
	if !(a == 55 && b == -55) {
		panic("err")
	}
}

func TestWithBigArray(t *testing.T) {

	a, b := WithBigArray()
	if !(a == 49995000 && b == -49995000) {
		panic("err")
	}
}

func TestWithoutBigArray(t *testing.T) {

	a, b := WithoutBigArray()
	if !(a == 49995000 && b == -49995000) {
		panic("err")
	}
}

func BenchmarkWithVariable(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		WithVariable()
	}
}

func BenchmarkWithoutVariable(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		WithoutVariable()
	}
}

func BenchmarkWithArray(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		WithArray()
	}
}

func BenchmarkWithoutArray(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		WithoutArray()
	}
}

func BenchmarkWithBigArray(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		WithBigArray()
	}
}

func BenchmarkWithoutBigArray(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		WithoutBigArray()
	}
}
