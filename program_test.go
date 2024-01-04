package formFormula_test

import (
	"fmt"
	"testing"

	formFormula "github.com/form-formula-go"
)

func Test_Disassemble(t *testing.T) {

	p := formFormula.NewProgram()

	p.NewOp(p.NewFunc((int)(formFormula.I), formFormula.FCT), p.NewFunc((int)(formFormula.E), formFormula.FCT), formFormula.SUM)

	//fmt.Printf("%v", p)

	fmt.Printf("%v", p.Disassemble())
}
