package formFormula_test

import (
	"fmt"
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
