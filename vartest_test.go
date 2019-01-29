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
