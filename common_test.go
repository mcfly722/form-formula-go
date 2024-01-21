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

func assert_int(t *testing.T, expected int, obtained int) {
	if obtained != expected {
		t.Fatalf("expected %v but obtained %v\n", expected, obtained)
	}
	t.Logf("obtained: %v\n", obtained)
}

func assert_error(t *testing.T, err error) {
	if err == nil {
		t.Fatalf("error not catched!")
	}
	t.Logf("successfully catched error: %v", err)
}
