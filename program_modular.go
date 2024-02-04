package formFormula

import (
	"errors"
	"fmt"
)

// Operations: +,-,*,pow

type ProgramModular interface {
	SetX(x uint64)
	NewFunc(operationType OperationType, operand1Offset uint) uint
	NewOp(operationType OperationType, operand1Offset uint, operand2Offset uint) uint
	Disassemble() string
	Dump() string
	Execute() uint64
	Recombine(maxXOccurrences uint, ready func())
	GetEstimation(maxXOccurrences uint) uint64
}

type programModular struct {
	memory            []uint64
	operations        []Operation
	byModule          uint64
	possibleConstants []uint
	possibleFunctions []uint
	possibleOperators []uint
}

func initializeMemoryForModularProgram(byModule uint64) []uint64 {
	memory := make([]uint64, len(Constants))
	memory[ONE] = 1
	memory[THREE] = 3
	memory[MINUS_ONE] = byModule - 1

	return memory
}

func newModularProgram(byModule uint64) *programModular {
	return &programModular{
		memory:     initializeMemoryForModularProgram(byModule),
		operations: []Operation{},
		byModule:   byModule,
		possibleConstants: []uint{
			uint(ONE),
			uint(MINUS_ONE),
			uint(THREE),
		},
		possibleFunctions: []uint{
			//			uint(FCT),
			uint(INVERSE),
		},
		possibleOperators: []uint{
			uint(SUM),
			uint(MUL),
			uint(POW),
			uint(GCD),
		},
	}
}

func NewModularProgram(byModule uint64) ProgramModular {
	return newModularProgram(byModule)
}

func NewModularProgramFromBracketsString(byModule uint64, bracketsString string) (ProgramModular, error) {
	if len(bracketsString) == 0 {
		return nil, errors.New("brackets sequence is empty")
	}

	tree, err := BracketsToExpressionTree(bracketsString)
	if err != nil {
		return nil, err
	}

	program := newModularProgram(byModule)

	_, err = program.loadFromExpressionTreeRecursive(tree)
	if err != nil {
		return nil, err
	}

	return program, nil
}

func (program *programModular) loadFromExpressionTreeRecursive(node *Expression) (uint, error) {
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

func (program *programModular) getPointersToFunctionsTypes() []*uint {
	result := []*uint{}
	l := len(program.operations)
	for i := 0; i < l; i++ {
		if program.operations[i].Operand2Offset == 0 {
			result = append(result, (*uint)(&program.operations[i].OperationType))
		}
	}
	return result
}

func (program *programModular) getPointersToOperatorsTypes() []*uint {
	result := []*uint{}
	l := len(program.operations)
	for i := 0; i < l; i++ {
		if program.operations[i].Operand2Offset != 0 {
			result = append(result, (*uint)(&program.operations[i].OperationType))
		}
	}
	return result
}

func (program *programModular) getPointersToConstantsOffsets() []*uint {
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

func (program *programModular) SetX(x uint64) {
	program.memory[X] = x
}

func (program *programModular) NewFunc(operationType OperationType, operand1Offset uint) uint {

	program.memory = append(program.memory, 666)
	resultAddr := len(program.memory) - 1

	newOp := Operation{
		Operand1Offset: uint(operand1Offset),
		OperationType:  operationType,
	}

	program.operations = append(program.operations, newOp)

	return uint(resultAddr)
}

func (program *programModular) NewOp(operationType OperationType, operand1Offset uint, operand2Offset uint) uint {
	program.memory = append(program.memory, 666)

	newOp := Operation{
		Operand1Offset: uint(operand1Offset),
		Operand2Offset: uint(operand2Offset),
		OperationType:  operationType,
	}
	program.operations = append(program.operations, newOp)
	return uint(len(program.memory) - 1)
}

func (program *programModular) toString(operation *Operation) string {
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

func (program *programModular) Disassemble() string {
	operation := program.operations[len(program.operations)-1]
	return fmt.Sprintf("(%v) mod %v", program.toString(&operation), program.byModule)
}

func (program *programModular) Dump() string {
	return fmt.Sprintf("memory:%v\nprogram:%v", program.memory, program.operations)
}

func pow2_uint64(x uint64, p uint64, m uint64) uint64 {
	if p == 0 {
		return x
	}
	for i := (uint64)(0); i < p; i++ {
		x = (x * x) % m
	}
	return x
}

// Pow returns x^n % m
func Internal_Pow_uint64_mod(x uint64, n uint64, m uint64) uint64 {
	c := n

	result := (uint64)(1)

	for j := (uint64)(0); c > 0; j++ {
		v := c % 2
		c = c / 2
		if v != 0 {
			result = (result * pow2_uint64(x, j, m)) % m
		}
	}

	return result
}

// Sub returns (m+a-b) % m
func Internal_Sub_uint64(a uint64, b uint64, m uint64) uint64 {
	return (m + a - b) % m
}

func Internal_Mul_uint64(a uint64, b uint64, m uint64) uint64 {
	return (a * b) % m
}

func Internal_GCD_uint64(a uint64, b uint64) uint64 {

	if a == 0 || b == 0 {
		return a + b
	}

	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}

func Internal_Inverse_uint64(a uint64, m uint64) uint64 {
	return Internal_Pow_uint64_mod(a, m-2, m)
}

func Internal_Add_uint64(a uint64, b uint64, m uint64) uint64 {
	return (a + b) % m
}

func (program *programModular) Execute() uint64 {

	memory := program.memory
	resultsOffset := len(Constants)

	for operationNumber, operation := range program.operations {
		memoryResultOffset := operationNumber + resultsOffset

		switch operation.OperationType {
		case SUM:
			memory[memoryResultOffset] = Internal_Add_uint64(memory[operation.Operand1Offset], memory[operation.Operand2Offset], program.byModule)
		case SUB:
			memory[memoryResultOffset] = Internal_Sub_uint64(memory[operation.Operand1Offset], memory[operation.Operand2Offset], program.byModule)
		case MUL:
			memory[memoryResultOffset] = Internal_Mul_uint64(memory[operation.Operand1Offset], memory[operation.Operand2Offset], program.byModule)
		case POW:
			memory[memoryResultOffset] = Internal_Pow_uint64_mod(memory[operation.Operand1Offset], memory[operation.Operand2Offset], program.byModule)
		case GCD:
			memory[memoryResultOffset] = Internal_GCD_uint64(memory[operation.Operand1Offset], memory[operation.Operand2Offset])
		case INVERSE:
			memory[memoryResultOffset] = Internal_Inverse_uint64(memory[operation.Operand1Offset], program.byModule)

		default:
			panic(fmt.Sprintf("unknown operationType=%v", operation.OperationType))
		}

		//fmt.Printf("%v %v %v -> %v\n", memory[operation.Operand1Offset], operations[operation.OperationType], memory[operation.Operand2Offset], memory[memoryResultOffset])
	}

	return program.memory[len(program.memory)-1]
}

// Recombine function
// ready(result) is the function which obtain calculation result, if this function returns
func (program *programModular) Recombine(maxXOccurrences uint, ready func()) {

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

func (program *programModular) GetEstimation(maxXOccurrences uint) uint64 {
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
