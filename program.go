package formFormula

import "fmt"

type OffsetMEM int

const (
	X   OffsetMEM = 0
	I   OffsetMEM = 1
	PVO OffsetMEM = 2
	PV1 OffsetMEM = 3
	PVX OffsetMEM = 4
	PI  OffsetMEM = 5
	E   OffsetMEM = 6
)

var Constants = map[OffsetMEM]string{
	X:   "x",
	I:   "i",
	PVO: "Pv0",
	PV1: "Pv1",
	PVX: "PvX",
	PI:  "Pi",
	E:   "e",
}

type OperationType int

const (
	SUM OperationType = 0
	MUL OperationType = 1
	DIV OperationType = 2
	FCT OperationType = 3
)

var operations = map[OperationType]string{
	SUM: "+",
	MUL: "*",
	DIV: "/",
	FCT: "!",
}

type Operation struct {
	Operand1Offset int
	Operand2Offset int
	OperationType  OperationType
}

type Program struct {
	memory     []float64
	operations []Operation
}

func NewProgram() *Program {
	return &Program{
		memory:     make([]float64, len(Constants)),
		operations: []Operation{},
	}
}

func (program *Program) NewFunc(operationType OperationType, operand1Offset int) int {

	program.memory = append(program.memory, 666)
	resultAddr := len(program.memory) - 1

	newOp := Operation{
		Operand1Offset: operand1Offset,
		OperationType:  operationType,
	}

	program.operations = append(program.operations, newOp)

	return resultAddr
}

func (program *Program) NewOp(operationType OperationType, operand1Offset int, operand2Offset int) int {
	program.memory = append(program.memory, 666)

	newOp := Operation{
		Operand1Offset: operand1Offset,
		Operand2Offset: operand2Offset,
		OperationType:  operationType,
	}
	program.operations = append(program.operations, newOp)
	return len(program.memory) - 1
}

func (program *Program) ToString(operation *Operation) string {
	result := ""

	if operation.Operand1Offset < len(Constants) {
		result += Constants[(OffsetMEM)(operation.Operand1Offset)]
	} else {
		op := program.operations[operation.Operand1Offset-len(Constants)]
		result += fmt.Sprintf("(%v)", program.ToString(&op))
	}

	result += operations[operation.OperationType]

	if operation.Operand2Offset != 0 {
		if operation.Operand2Offset < len(Constants) {
			result += Constants[(OffsetMEM)(operation.Operand2Offset)]
		} else {
			op := program.operations[operation.Operand2Offset-len(Constants)]
			result += fmt.Sprintf("(%v)", program.ToString(&op))
		}
	}

	return result
}

func (program *Program) Disassemble() string {
	operation := program.operations[len(program.operations)-1]
	return program.ToString(&operation)
}
