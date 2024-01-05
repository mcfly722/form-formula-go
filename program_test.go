package formFormula_test

import (
	"fmt"
	"math"
	"testing"

	formFormula "github.com/form-formula-go"
)

func Test_Disassemble(t *testing.T) {

	p := formFormula.NewProgram()

	p.NewOp(
		formFormula.SUM,
		p.NewFunc(formFormula.FCT, (int)(formFormula.I)),
		p.NewOp(formFormula.MUL,
			int(formFormula.X),
			p.NewFunc(formFormula.FCT, (int)(formFormula.E)),
		),
	)

	disassembled := p.Disassemble()

	fmt.Printf("%v", disassembled)

	if disassembled != "(i!)+(x*(e!))" {
		t.Fatal("not expected disassembled code")
	}
}

func Test_Execute(t *testing.T) {
	p := formFormula.NewProgram()

	p.SetX(2)

	p.NewOp(
		formFormula.MUL,
		int(formFormula.E),
		p.NewOp(formFormula.SUM,
			int(formFormula.X),
			int(formFormula.PI),
		),
	)

	fmt.Printf("program: %v\n", p.Disassemble())

	result := p.Execute()
	fmt.Printf("dump:\n%v\n", p.Dump())
	fmt.Printf("result: %v\n", result)

	expected := math.E * (2 + math.Pi)
	if result != expected {
		t.Fatalf("expected=%v result=%v", result, expected)
	}
}

func Test_Execute_Empty(t *testing.T) {
	p := formFormula.NewProgram()
	p.Execute()
}
