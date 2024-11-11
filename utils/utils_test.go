package utils

import (
	"testing"
)

func TestIfInt(t *testing.T) {
	result := If(true, 42, 0)
	if result != 42 {
		t.Errorf("If(true, 42, 0) = %d; want 42", result)
	}

	result = If(false, 42, 0)
	if result != 0 {
		t.Errorf("If(false, 42, 0) = %d; want 0", result)
	}
}

func TestIfString(t *testing.T) {
	result := If(true, "yes", "no")
	if result != "yes" {
		t.Errorf(`If(true, "yes", "no") = %s; want "yes"`, result)
	}

	result = If(false, "yes", "no")
	if result != "no" {
		t.Errorf(`If(false, "yes", "no") = %s; want "no"`, result)
	}
}

func TestIfBool(t *testing.T) {
	result := If(true, true, false)
	if result != true {
		t.Errorf("If(true, true, false) = %v; want true", result)
	}

	result = If(false, true, false)
	if result != false {
		t.Errorf("If(false, true, false) = %v; want false", result)
	}
}
