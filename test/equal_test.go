package test

import (
	"testing"
)

func TestEqual(t *testing.T) {
	a := 1
	b := 1
	shouldBe := true
	if real := equal(a, b); real != shouldBe {
		t.Errorf("equal(%d, %d) should be %v, but is:%v\n", a, b, shouldBe, real)
	}
}
