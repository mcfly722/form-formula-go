package formFormula_test

import (
	"math"
	"testing"

	formFormula "github.com/form-formula-go"
)

func Test_ProgramIterational_Disassemble(t *testing.T) {

	p := formFormula.NewIterationalProgram()

	p.NewOp(
		formFormula.SUM,
		p.NewFunc(formFormula.FCT, uint(formFormula.I)),
		p.NewOp(formFormula.MUL,
			uint(formFormula.X),
			p.NewFunc(formFormula.FCT, uint(formFormula.E)),
		),
	)

	assert_string(t, "(i!)+(x*(e!))", p.Disassemble())
}

func Test_ProgramIterational_Execute(t *testing.T) {
	p := formFormula.NewIterationalProgram()

	p.SetX(2)

	p.NewOp(
		formFormula.MUL,
		uint(formFormula.E),
		p.NewOp(formFormula.SUM,
			uint(formFormula.X),
			uint(formFormula.PI),
		),
	)

	assert_float64(t, math.E*(2+math.Pi), p.Execute())
}

func Test_ProgramIterational_Execute_Empty(t *testing.T) {
	p := formFormula.NewIterationalProgram()
	p.Execute()
}

func Test_ProgramIterational_ExecuteWithIterations_I(t *testing.T) {
	p := formFormula.NewIterationalProgram()

	p.NewOp(
		formFormula.SUM,
		uint(formFormula.PVX),
		uint(formFormula.I),
	)

	// 3+1=4
	// 4+2=6
	// 6+3=9

	assert_float64(t, 9, p.ExecuteWithIterations(3, 3))
}

func Test_ProgramIterational_ExecuteWithIterations_PVX(t *testing.T) {
	p := formFormula.NewIterationalProgram()

	p.NewOp(
		formFormula.SUM,
		uint(formFormula.PVX),
		uint(formFormula.MINUS_ONE),
	)

	// 10-1=9
	//  9-1=8
	//  8-1=7

	assert_string(t, "PvX+-1", p.Disassemble())
	assert_float64(t, 7, p.ExecuteWithIterations(3, 10))
}

func Test_ProgramIterational_ExecuteWithIterations_PV0(t *testing.T) {
	p := formFormula.NewIterationalProgram()

	p.NewOp(
		formFormula.SUM,
		uint(formFormula.PV0),
		uint(formFormula.X),
	)

	//  0-2=-2
	// -2-2=-4
	// -4-2=-6

	assert_float64(t, -6, p.ExecuteWithIterations(3, -2))
}

func Test_IterationalProgram_UnknownOperationType(t *testing.T) {
	p := formFormula.NewIterationalProgram()

	p.NewOp(
		666,
		uint(formFormula.ONE),
		uint(formFormula.ONE),
	)

	defer func(t *testing.T) {
		if err := recover(); err != nil {
			t.Logf("panic successfully catched: %v", err)
		}
	}(t)

	p.Execute()

	t.Fatal("panic not catched!")
}

func Test_IterationalProgram_Dump(t *testing.T) {
	p, err := formFormula.NewIterationalProgramFromBracketsString("(()())(())")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("DUMP:\n%v", p.Dump())
}

func Test_IterationalProgram_Empty(t *testing.T) {
	_, err := formFormula.NewIterationalProgramFromBracketsString("")
	assert_error(t, err)
}

func Test_IterationalProgram_NewIterationalProgramFromBracketsString_Error(t *testing.T) {
	_, err := formFormula.NewIterationalProgramFromBracketsString("(()")
	assert_error(t, err)
}

func Test_IterationalProgram_ThreeArguments_Error(t *testing.T) {
	_, err := formFormula.NewIterationalProgramFromBracketsString("()()()")
	assert_error(t, err)
}

func Test_IterationalProgram_ThreeArguments_ForFunction_Error(t *testing.T) {
	_, err := formFormula.NewIterationalProgramFromBracketsString("(()()())")
	assert_error(t, err)
}

func Test_IterationalProgram_ThreeArguments_ForFirstOperand_Error(t *testing.T) {
	_, err := formFormula.NewIterationalProgramFromBracketsString("(()()())()")
	assert_error(t, err)
}

func Test_IterationalProgram_ThreeArguments_ForSecondOperand_Error(t *testing.T) {
	_, err := formFormula.NewIterationalProgramFromBracketsString("()(()()())")
	assert_error(t, err)
}

func Test_ProgramIterational_Infinity(t *testing.T) {
	p := formFormula.NewIterationalProgram()

	p.NewOp(
		formFormula.DIV,
		uint(formFormula.ONE),
		uint(formFormula.X),
	)

	assert_string(t, "1/x", p.Disassemble())

	value := p.ExecuteWithIterations(1, 0)
	assert_bool(t, math.IsInf(value, 1))
}

func Test_ProgramIterational_Nothing(t *testing.T) {
	p := formFormula.NewIterationalProgram()

	p.NewFunc(
		formFormula.NOTHING,
		uint(formFormula.X),
	)

	assert_string(t, "nothing(x)", p.Disassemble())

	assert_float64(t, 1234567, p.ExecuteWithIterations(1, 1234567))
}
