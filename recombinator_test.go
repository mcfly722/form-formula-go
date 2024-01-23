package formFormula_test

import (
	"fmt"
	"testing"

	formFormula "github.com/form-formula-go"
)

func setLadder(input *[]*uint) {
	for i := uint(0); i < uint(len(*input)); i++ {
		*(*input)[i] = i + 1
	}
}
func Test_SetLadder(t *testing.T) {
	input := [5]uint{}
	inputRefs := []*uint{&input[0], &input[1], &input[2], &input[3], &input[4]}
	setLadder(&inputRefs)
	assert_string(t, "[1 2 3 4 5]", fmt.Sprintf("%v", input))
}

func Test_SetLadderEmpty(t *testing.T) {
	input := []uint{}
	inputRefs := []*uint{}
	setLadder(&inputRefs)
	assert_string(t, "[]", fmt.Sprintf("%v", input))
}

func Test_RecombineRequiredX(t *testing.T) {
	input := [5]uint{}
	inputRefs := []*uint{&input[0], &input[1], &input[2], &input[3], &input[4]}

	counter := 0

	next := func(remaining *[]*uint) {
		counter++
		setLadder(remaining)
		fmt.Printf("%2v %v\n", counter, input)
	}

	formFormula.RecombineRequiredX(&inputRefs, 5, 0, next)

	assert_int(t, 31, counter)
}

func Test_RecombineRequiredX_EmptyInput(t *testing.T) {
	input := []uint{}
	inputRefs := []*uint{}

	counter := 0

	next := func(remaining *[]*uint) {
		counter++
		setLadder(remaining)
		fmt.Printf("%2v %v\n", counter, input)
	}

	formFormula.RecombineRequiredX(&inputRefs, 5, 0, next)

	assert_int(t, 1, counter)
}

func Test_Recombine(t *testing.T) {
	input := [5]uint{}
	inputRefs := []*uint{&input[0], &input[1], &input[2], &input[3], &input[4]}

	counter := 0

	ready := func() {
		counter++
		fmt.Printf("%3v %v\n", counter, input)
	}

	formFormula.RecombineValues(&inputRefs, &[]uint{0, 1, 2}, ready)

	assert_int(t, 3*3*3*3*3, counter)
}

func Test_Recombine_EmptyInput(t *testing.T) {
	input := []uint{}
	inputRefs := []*uint{}

	counter := 0

	ready := func() {
		counter++
		fmt.Printf("%3v %v\n", counter, input)
	}

	formFormula.RecombineValues(&inputRefs, &[]uint{0, 1, 2}, ready)

	assert_int(t, 1, counter)
}

func Test_Recombine_EmptyValues(t *testing.T) {
	input := [5]uint{}
	inputRefs := []*uint{&input[0], &input[1], &input[2], &input[3], &input[4]}

	counter := 0

	ready := func() {
		counter++
		fmt.Printf("%3v %v\n", counter, input)
	}

	formFormula.RecombineValues(&inputRefs, &[]uint{}, ready)

	assert_int(t, 1, counter)
}
