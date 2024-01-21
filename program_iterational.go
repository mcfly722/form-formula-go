package formFormula

import (
	"errors"
	"fmt"
	"math"
)

type ProgramIterational interface {
	SetX(x float64)
	NewFunc(operationType OperationType, operand1Offset int) int
	NewOp(operationType OperationType, operand1Offset int, operand2Offset int) int
	Disassemble() string
	Dump() string
	Execute() float64
	ExecuteWithIterations(n int, x float64) float64
}

type programIterational struct {
	memory            []float64
	operations        []Operation
	possibleConstants []int
	possibleFunctions []int
	possibleOperators []int
}

func NewIterationalProgram() ProgramIterational {
	memory := make([]float64, len(Constants))

	memory[ONE] = 1
	memory[MINUS_ONE] = -1
	memory[PI] = math.Pi
	memory[E] = math.E

	return &programIterational{
		memory:     memory,
		operations: []Operation{},
	}
}

func initializeMemoryForIterationalProgram() []float64 {
	memory := make([]float64, len(Constants))
	memory[ONE] = 1
	memory[THREE] = 3
	return memory
}

func newIterationalProgram() *programIterational {
	return &programIterational{
		memory:     initializeMemoryForIterationalProgram(),
		operations: []Operation{},
		possibleConstants: []int{
			int(ONE),
			int(MINUS_ONE),
			int(THREE),
		},
		possibleFunctions: []int{
			int(FCT),
			int(INVERSE),
		},
		possibleOperators: []int{
			int(SUM),
			int(MUL),
			int(POW),
		},
	}
}

func NewIterationalProgramFromBracketsString(bracketsString string) (ProgramIterational, error) {
	if len(bracketsString) == 0 {
		return nil, errors.New("brackets sequence is empty")
	}

	tree, err := BracketsToExpressionTree(bracketsString)
	if err != nil {
		return nil, err
	}

	program := newIterationalProgram()

	_, err = program.loadFromExpressionTreeRecursive(tree)
	if err != nil {
		return nil, err
	}

	return program, nil
}

func (program *programIterational) loadFromExpressionTreeRecursive(node *Expression) (int, error) {
	switch len(node.Arguments) {
	case 0:
		return int(X), nil
	case 1:
		argumentOffset, err := program.loadFromExpressionTreeRecursive(node.Arguments[0])
		if err != nil {
			return -1, err
		}
		operation := Operation{
			Operand1Offset: argumentOffset,
			OperationType:  FCT,
		}
		program.memory = append(program.memory, 0)
		program.operations = append(program.operations, operation)
		return len(program.memory) - 1, nil
	case 2:
		argumentOffset1, err := program.loadFromExpressionTreeRecursive(node.Arguments[0])
		if err != nil {
			return -1, err
		}
		argumentOffset2, err := program.loadFromExpressionTreeRecursive(node.Arguments[1])
		if err != nil {
			return -1, err
		}
		operation := Operation{
			Operand1Offset: argumentOffset1,
			Operand2Offset: argumentOffset2,
			OperationType:  SUM,
		}
		program.memory = append(program.memory, 0)
		program.operations = append(program.operations, operation)
		return len(program.memory) - 1, nil
	default:
		return -1, errors.New("three arguments not supported by modular arithmetic")
	}
}

func (program *programIterational) SetX(x float64) {
	program.memory[X] = x
}

func (program *programIterational) NewFunc(operationType OperationType, operand1Offset int) int {

	program.memory = append(program.memory, 666)
	resultAddr := len(program.memory) - 1

	newOp := Operation{
		Operand1Offset: operand1Offset,
		OperationType:  operationType,
	}

	program.operations = append(program.operations, newOp)

	return resultAddr
}

func (program *programIterational) NewOp(operationType OperationType, operand1Offset int, operand2Offset int) int {
	program.memory = append(program.memory, 666)

	newOp := Operation{
		Operand1Offset: operand1Offset,
		Operand2Offset: operand2Offset,
		OperationType:  operationType,
	}
	program.operations = append(program.operations, newOp)
	return len(program.memory) - 1
}

func (program *programIterational) toString(operation *Operation) string {
	val1 := ""
	val2 := ""

	if operation.Operand1Offset < len(Constants) {
		val1 = Constants[(OffsetMEM)(operation.Operand1Offset)]
	} else {
		op := program.operations[operation.Operand1Offset-len(Constants)]
		val1 += fmt.Sprintf("(%v)", program.toString(&op))
	}

	if operation.Operand2Offset != 0 {
		if operation.Operand2Offset < len(Constants) {
			val2 = Constants[(OffsetMEM)(operation.Operand2Offset)]
		} else {
			op := program.operations[operation.Operand2Offset-len(Constants)]
			val2 = fmt.Sprintf("(%v)", program.toString(&op))
		}
	}

	return operations[operation.OperationType](val1, val2)
}

func (program *programIterational) Disassemble() string {
	operation := program.operations[len(program.operations)-1]
	return program.toString(&operation)
}

func (program *programIterational) Dump() string {
	return fmt.Sprintf("memory:%v\nprogram:%v", program.memory, program.operations)
}

func (program *programIterational) Execute() float64 {

	memory := program.memory
	resultsOffset := len(Constants)

	for operationNumber, operation := range program.operations {
		memoryResultOffset := operationNumber + resultsOffset

		switch operation.OperationType {
		case NOTHING:
			memory[memoryResultOffset] = memory[operation.Operand1Offset]
		case SUM:
			memory[memoryResultOffset] = memory[operation.Operand1Offset] + memory[operation.Operand2Offset]
		case MUL:
			memory[memoryResultOffset] = memory[operation.Operand1Offset] * memory[operation.Operand2Offset]
		case DIV:
			memory[memoryResultOffset] = memory[operation.Operand1Offset] / memory[operation.Operand2Offset]
		default:
			panic(fmt.Sprintf("unknown operationType=%v", operation.OperationType))
		}
	}

	return program.memory[len(program.memory)-1]
}

func (program *programIterational) ExecuteWithIterations(n int, x float64) float64 {

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
