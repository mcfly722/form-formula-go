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
	fmt.Printf("result=%v\n", result)

	expected := math.E * (2 + math.Pi)
	if result != expected {
		t.Fatalf("expected=%v", expected)
	}
}

func Test_Execute_Empty(t *testing.T) {
	p := formFormula.NewProgram()
	p.Execute()
}

func Test_ExecuteWithIterations_I(t *testing.T) {
	p := formFormula.NewProgram()

	p.NewOp(
		formFormula.SUM,
		int(formFormula.PVX),
		int(formFormula.I),
	)

	// 3+1=4
	// 4+2=6
	// 6+3=9

	result := p.ExecuteWithIterations(3, 3)
	fmt.Printf("result=%v\n", result)
	expected := float64(9)

	if result != expected {
		t.Fatalf("expected=%v", expected)
	}
}

func Test_ExecuteWithIterations_PVX(t *testing.T) {
	p := formFormula.NewProgram()

	p.NewOp(
		formFormula.SUM,
		int(formFormula.PVX),
		int(formFormula.MINUS_ONE),
	)

	// 10-1=9
	// 9-1=8
	// 8-1=7

	result := p.ExecuteWithIterations(3, 10)
	fmt.Printf("result=%v\n", result)
	expected := float64(7)

	if result != expected {
		t.Fatalf("expected=%v", expected)
	}
}

func Test_ExecuteWithIterations_PV0(t *testing.T) {
	p := formFormula.NewProgram()

	p.NewOp(
		formFormula.SUM,
		int(formFormula.PV0),
		int(formFormula.X),
	)

	// 0-2=-2
	// -2-2=-4
	// -4-2=-6

	result := p.ExecuteWithIterations(3, -2)
	fmt.Printf("result=%v\n", result)
	expected := float64(-6)

	if result != expected {
		t.Fatalf("expected=%v", expected)
	}
}
