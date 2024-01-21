package formFormula_test

import (
	"testing"

	formFormula "github.com/form-formula-go"
)

func Test_Pow_uint64(t *testing.T) {
	assert_uint64(t, 1594323, formFormula.Internal_Pow_uint64(3, 13))
}

func Test_Fact_uint64(t *testing.T) {
	assert_uint64(t, 1307674368000, formFormula.Internal_Fact_uint64(15))
}

func Test_Combination_uint64(t *testing.T) {
	assert_uint64(t, 10, formFormula.Internal_Combination_uint64(5, 3))
}

func Test_GetEstimation(t *testing.T) {
	assert_uint64(t, 1575936, formFormula.Internal_GetEstimation(
		3,
		2,
		3,
		4,
		3,
		4,
		5,
	))
}
