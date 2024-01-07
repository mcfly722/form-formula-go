package formFormula_test

import (
	"testing"
)

func assert_string(t *testing.T, expected string, obtained string) {
	if obtained != expected {
		t.Fatalf("expected %v but obtained %v\n", expected, obtained)
	}
	t.Logf("obtained: %v\n", obtained)
}

func assert_float64(t *testing.T, expected float64, obtained float64) {
	if obtained != expected {
		t.Fatalf("expected %v but obtained %v\n", expected, obtained)
	}
	t.Logf("obtained: %v\n", obtained)
}

func assert_uint64(t *testing.T, expected uint64, obtained uint64) {
	if obtained != expected {
		t.Fatalf("expected %v but obtained %v\n", expected, obtained)
	}
	t.Logf("obtained: %v\n", obtained)
}
