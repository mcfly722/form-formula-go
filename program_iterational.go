package formFormula

import (
	"errors"
	"fmt"
	"math"
)

type ProgramIterational interface {
	SetX(x float64)
	NewFunc(operationType OperationType, operand1Offset uint) uint
	NewOp(operationType OperationType, operand1Offset uint, operand2Offset uint) uint
	Disassemble() string
	Dump() string
	ExecuteWithIterations(iterations uint, x float64) float64
	RecombineForms(maxXOccurrences uint, ready func())
	GetEstimation(maxXOccurrences uint) uint64
}

type programIterational struct {
	memory            []float64
	operations        []Operation
	possibleConstants []uint
	possibleFunctions []uint
	possibleOperators []uint
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
		possibleConstants: []uint{
			uint(ONE),
			uint(MINUS_ONE),
			uint(THREE),
		},
		possibleFunctions: []uint{
			uint(FCT),
			uint(SQRT),
		},
		possibleOperators: []uint{
			uint(SUM),
			uint(MUL),
			uint(POW),
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

func (program *programIterational) loadFromExpressionTreeRecursive(node *Expression) (uint, error) {
	switch len(node.Arguments) {
	case 0:
		return uint(X), nil
	case 1:
		argumentOffset, err := program.loadFromExpressionTreeRecursive(node.Arguments[0])
		if err != nil {
			return defaultErrorUIntValue, err
		}
		operation := Operation{
			Operand1Offset: argumentOffset,
			OperationType:  FCT,
		}
		program.memory = append(program.memory, 0)
		program.operations = append(program.operations, operation)
		return uint(len(program.memory) - 1), nil
	case 2:
		argumentOffset1, err := program.loadFromExpressionTreeRecursive(node.Arguments[0])
		if err != nil {
			return defaultErrorUIntValue, err
		}
		argumentOffset2, err := program.loadFromExpressionTreeRecursive(node.Arguments[1])
		if err != nil {
			return defaultErrorUIntValue, err
		}
		operation := Operation{
			Operand1Offset: argumentOffset1,
			Operand2Offset: argumentOffset2,
			OperationType:  SUM,
		}
		program.memory = append(program.memory, 0)
		program.operations = append(program.operations, operation)
		return uint(len(program.memory) - 1), nil
	default:
		return defaultErrorUIntValue, errors.New("three arguments not supported by modular arithmetic")
	}
}

func (program *programIterational) SetX(x float64) {
	program.memory[X] = x
}

func (program *programIterational) NewFunc(operationType OperationType, operand1Offset uint) uint {

	program.memory = append(program.memory, 666)
	resultAddr := uint(len(program.memory) - 1)

	newOp := Operation{
		Operand1Offset: operand1Offset,
		OperationType:  operationType,
	}

	program.operations = append(program.operations, newOp)

	return resultAddr
}

func (program *programIterational) NewOp(operationType OperationType, operand1Offset uint, operand2Offset uint) uint {
	program.memory = append(program.memory, 666)

	newOp := Operation{
		Operand1Offset: operand1Offset,
		Operand2Offset: operand2Offset,
		OperationType:  operationType,
	}
	program.operations = append(program.operations, newOp)
	return uint(len(program.memory) - 1)
}

func (program *programIterational) toString(operation *Operation) string {
	val1 := ""
	val2 := ""

	if operation.Operand1Offset < uint(len(Constants)) {
		val1 = Constants[(OffsetMEM)(operation.Operand1Offset)]
	} else {
		op := program.operations[int(operation.Operand1Offset)-len(Constants)]
		val1 += fmt.Sprintf("(%v)", program.toString(&op))
	}

	if operation.Operand2Offset != 0 {
		if operation.Operand2Offset < uint(len(Constants)) {
			val2 = Constants[(OffsetMEM)(operation.Operand2Offset)]
		} else {
			op := program.operations[int(operation.Operand2Offset)-len(Constants)]
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

func (program *programIterational) execute() float64 {

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

func (program *programIterational) ExecuteWithIterations(iterations uint, x float64) float64 {

	program.memory[PV0] = 0
	program.memory[PV1] = 1
	program.memory[PVX] = x
	program.memory[X] = x

	resultAddr := len(program.memory) - 1

	for i := uint(1); i <= iterations; i++ {
		program.memory[I] = float64(i)
		result := program.execute()

		if math.IsNaN(result) || math.IsInf(result, 1) || math.IsInf(result, -1) {
			break
		}

		program.memory[PV0] = result
		program.memory[PV1] = result
		program.memory[PVX] = result
	}

	return program.memory[resultAddr]
}

func (program *programIterational) getPointersToFunctionsTypes() []*uint {
	result := []*uint{}
	l := len(program.operations)
	for i := 0; i < l; i++ {
		if program.operations[i].Operand2Offset == 0 {
			result = append(result, (*uint)(&program.operations[i].OperationType))
		}
	}
	return result
}

func (program *programIterational) getPointersToOperatorsTypes() []*uint {
	result := []*uint{}
	l := len(program.operations)
	for i := 0; i < l; i++ {
		if program.operations[i].Operand2Offset != 0 {
			result = append(result, (*uint)(&program.operations[i].OperationType))
		}
	}
	return result
}

func (program *programIterational) getPointersToConstantsOffsets() []*uint {
	constantsRange := uint(len(Constants))
	result := []*uint{}

	l := len(program.operations)

	for i := 0; i < l; i++ {
		if program.operations[i].Operand1Offset < constantsRange {
			//fmt.Printf("OP1:%v [%v]\n", Constants[(OffsetMEM)(program.operations[i].Operand1Offset)], &(program.operations[i].Operand2Offset))
			result = append(result, &(program.operations[i].Operand1Offset))
		}

		if program.operations[i].Operand2Offset != 0 && program.operations[i].Operand2Offset < constantsRange {
			//fmt.Printf("OP2:%v [%v]\n", Constants[(OffsetMEM)(program.operations[i].Operand2Offset)], &(program.operations[i].Operand2Offset))
			result = append(result, &(program.operations[i].Operand2Offset))
		}
	}
	return result
}

func (program *programIterational) RecombineForms(maxXOccurrences uint, ready func()) {
	constants := program.getPointersToConstantsOffsets()
	functions := program.getPointersToFunctionsTypes()
	operations := program.getPointersToOperatorsTypes()

	ready_X_Constants_Functions := func() {
		RecombineValues(&operations, &program.possibleOperators, ready)
	}

	ready_X_Constants := func() {
		RecombineValues(&functions, &program.possibleFunctions, ready_X_Constants_Functions)
	}

	readyX := func(remainedConstants *[]*uint) {
		RecombineValues(remainedConstants, &program.possibleConstants, ready_X_Constants)
	}

	RecombineRequiredX(&constants, maxXOccurrences, uint(X), readyX)
}

func (program *programIterational) GetEstimation(maxXOccurrences uint) uint64 {
	return Internal_GetEstimation(
		maxXOccurrences,
		uint(len(program.possibleConstants)),
		uint(len(program.possibleFunctions)),
		uint(len(program.possibleOperators)),
		uint(len(program.getPointersToConstantsOffsets())),
		uint(len(program.getPointersToFunctionsTypes())),
		uint(len(program.getPointersToOperatorsTypes())),
	)
}
