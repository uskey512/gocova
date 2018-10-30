package main

import "testing"

func TestClamp(t *testing.T) {
	if Clamp(5.0, 0.0, 10.0) != 5.0 {
		t.Fail()
	}
	if Clamp(-5.0, 0.0, 10.0) != 0.0 {
		t.Fail()
	}
	if Clamp(15.0, 0.0, 10.0) != 10.0 {
		t.Fail()
	}
}

func TestClamp01(t *testing.T) {
	if Clamp01(0.5) != 0.5 {
		t.Fail()
	}
	if Clamp01(-0.5) != 0.0 {
		t.Fail()
	}
	if Clamp01(2.0) != 1.0 {
		t.Fail()
	}
}
