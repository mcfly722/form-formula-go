package formFormula_test

import (
	"fmt"
	"testing"

	formFormula "github.com/form-formula-go"
)

func defaultTestModularProgram() formFormula.ProgramModular {

	p := formFormula.NewModularProgram(29)

	p.NewOp(
		formFormula.SUM,
		(int)(formFormula.ONE),
		p.NewOp(formFormula.MUL,
			int(formFormula.X),
			p.NewOp(formFormula.POW,
				(int)(formFormula.THREE),
				p.NewFunc(formFormula.FCT, (int)(formFormula.X)),
			),
		),
	)
	return p
}

func Test_ProgramModular_Disassemble(t *testing.T) {

	p := formFormula.NewModularProgram(4)

	p.NewOp(
		formFormula.SUM,
		(int)(formFormula.ONE),
		p.NewOp(formFormula.MUL,
			int(formFormula.X),
			p.NewOp(formFormula.POW,
				(int)(formFormula.THREE),
				(int)(formFormula.X),
			),
		),
	)

	assert_string(t, "(1+(x*(3^x))) mod 4", p.Disassemble())
}

func Test_Sub_uint64_1(t *testing.T) {
	assert_uint64(t, 2, formFormula.Sub_uint64(3, 5, 4))
}

func Test_Sub_uint64_2(t *testing.T) {
	assert_uint64(t, 0, formFormula.Sub_uint64(3, 3, 4))
}

func Test_Pow_uint64_zero(t *testing.T) {
	assert_uint64(t, 1, formFormula.Pow_uint64(3453453453, 0, 2342341))
}

func Test_Pow_uint64_one(t *testing.T) {
	assert_uint64(t, 3453453453%2342341, formFormula.Pow_uint64(3453453453, 1, 2342341))
}

func Test_Pow_uint64_2326182(t *testing.T) {
	assert_uint64(t, 2326182, formFormula.Pow_uint64(3453453453, 437483784, 2342341))
}

func Test_GCD_uint64_primes(t *testing.T) {
	assert_uint64(t, 1, formFormula.GCD_uint64(2867395040399, 6816348081737))
}

func Test_GCD_uint64(t *testing.T) {
	prime := (uint64)(634741)
	assert_uint64(t, prime, formFormula.GCD_uint64(prime*233837, prime*975643))
}

func Test_Execute(t *testing.T) {

	p := formFormula.NewModularProgram(123902934)
	p.SetX(7)

	p.NewOp(
		formFormula.SUM,
		(int)(formFormula.ONE),
		p.NewOp(formFormula.MUL,
			int(formFormula.X),
			p.NewOp(formFormula.POW,
				(int)(formFormula.THREE),
				(int)(formFormula.X),
			),
		),
	)

	fmt.Printf("program:%v\n", p.Disassemble())

	assert_uint64(t, 15310, p.Execute())
}

func Test_ChangeOperators(t *testing.T) {
	p := defaultTestModularProgram()
	fmt.Printf("%v\n", p.Disassemble())
	operatorsAddresses := p.GetPointersToOperatorsTypes()
	for _, operator := range operatorsAddresses {
		*operator = int(formFormula.SUB)
	}
	assert_string(t, "(1-(x-(3-(x!)))) mod 29", p.Disassemble())
}

func Test_ChangeConstants(t *testing.T) {
	p := defaultTestModularProgram()
	fmt.Printf("%v\n", p.Disassemble())
	constantsPointers := p.GetPointersToConstantsOffsets()
	for _, constantPointer := range constantsPointers {
		*constantPointer = int(formFormula.X)
	}
	assert_string(t, "(x+(x*(x^(x!)))) mod 29", p.Disassemble())
}

func Test_ChangeFunctions(t *testing.T) {
	p := defaultTestModularProgram()
	fmt.Printf("%v\n", p.Disassemble())
	operatorsAddresses := p.GetPointersToFunctionsTypes()
	for _, operator := range operatorsAddresses {
		*operator = int(formFormula.INVERSE)
	}
	assert_string(t, "(1+(x*(3^(inverse(x))))) mod 29", p.Disassemble())
}

func Test_NewModularProgramFromBracketsString(t *testing.T) {
	p, err := formFormula.NewModularProgramFromBracketsString(15, "(()())((()))")
	if err != nil {
		t.Fatal(err)
	}
	assert_string(t, "((x+x)+((x!)!)) mod 15", p.Disassemble())
}

func Test_NewModularProgramFromBracketsString_Empty(t *testing.T) {
	_, err := formFormula.NewModularProgramFromBracketsString(15, "")
	if err == nil {
		t.Fatal("error for empty bracket sequence is not catched!")
	}
	fmt.Printf("Successfully catched error = %v\n", err)
}

func Test_NewModularProgramFromBracketsString_TwoBracketsPairs(t *testing.T) {
	p, err := formFormula.NewModularProgramFromBracketsString(15, "()()")
	if err != nil {
		t.Fatal(err)
	}
	assert_string(t, "(x+x) mod 15", p.Disassemble())
}

func Test_NewModularProgramFromBracketsString_OneBracketsPair(t *testing.T) {
	p, err := formFormula.NewModularProgramFromBracketsString(15, "()")
	if err != nil {
		t.Fatal(err)
	}

	// () is not an constant, it is operation over constant
	assert_string(t, "(x!) mod 15", p.Disassemble())
}

func Test_NewModularProgramFromBracketsString_OneBracketsPair2(t *testing.T) {
	p, err := formFormula.NewModularProgramFromBracketsString(15, "(())")
	if err != nil {
		t.Fatal(err)
	}

	// () is not an constant, it is operation over constant
	assert_string(t, "((x!)!) mod 15", p.Disassemble())
}
