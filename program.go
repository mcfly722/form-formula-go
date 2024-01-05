package formFormula

import (
	"fmt"
	"math"
)

type OffsetMEM int

const (
	X         OffsetMEM = 0
	I         OffsetMEM = 1
	PV0       OffsetMEM = 2
	PV1       OffsetMEM = 3
	PVX       OffsetMEM = 4
	ONE       OffsetMEM = 5
	MINUS_ONE OffsetMEM = 6
	PI        OffsetMEM = 7
	E         OffsetMEM = 8
)

var Constants = map[OffsetMEM]string{
	X:         "x",
	I:         "i",
	PV0:       "Pv0",
	PV1:       "Pv1",
	PVX:       "PvX",
	ONE:       "1",
	MINUS_ONE: "-1",
	PI:        "Pi",
	E:         "e",
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
	memory := make([]float64, len(Constants))

	memory[ONE] = 1
	memory[MINUS_ONE] = -1
	memory[PI] = math.Pi
	memory[E] = math.E

	return &Program{
		memory:     memory,
		operations: []Operation{},
	}
}

func (program *Program) SetX(x float64) {
	program.memory[X] = x
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

func (program *Program) Dump() string {
	return fmt.Sprintf("memory:%v\nprogram:%v", program.memory, program.operations)
}

func (program *Program) Execute() float64 {

	memory := program.memory
	resultsOffset := len(Constants)

	for operationNumber, operation := range program.operations {
		memoryResultOffset := operationNumber + resultsOffset

		switch operation.OperationType {
		case SUM:
			memory[memoryResultOffset] = memory[operation.Operand1Offset] + memory[operation.Operand2Offset]
		case MUL:
			memory[memoryResultOffset] = memory[operation.Operand1Offset] * memory[operation.Operand2Offset]
		default:
			panic(fmt.Sprintf("unknown operationType=%v", operation.OperationType))
		}
	}

	return program.memory[len(program.memory)-1]
}

func (program *Program) ExecuteWithIterations(n int, x float64) float64 {

	program.memory[PV0] = 0
	program.memory[PV1] = 1
	program.memory[PVX] = x
	program.memory[X] = x

	resultAddr := len(program.memory) - 1

	for i := 1; i <= n; i++ {
		program.memory[I] = float64(i)
		result := program.Execute()

		if math.IsNaN(result) || math.IsInf(result, 1) || math.IsInf(result, -1) {
			break
		}

		program.memory[PV0] = result
		program.memory[PV1] = result
		program.memory[PVX] = result
	}

	return program.memory[resultAddr]
}
