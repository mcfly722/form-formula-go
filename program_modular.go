package formFormula

import (
	"errors"
	"fmt"
)

// Operations: +,-,*,pow

type ProgramModular interface {
	SetX(x uint64)
	NewFunc(operationType OperationType, operand1Offset int) int
	NewOp(operationType OperationType, operand1Offset int, operand2Offset int) int
	Disassemble() string
	Dump() string
	Execute() uint64
	GetPointersToFunctionsTypes() []*int
	GetPointersToOperatorsTypes() []*int
	GetPointersToConstantsOffsets() []*int
	Recombine(x []uint64, maxXOccurrences int, ready func())
	GetEstimation(maxXOccurrences int) uint64
}

type programModular struct {
	memory            []uint64
	operations        []Operation
	byModule          uint64
	possibleConstants []int
	possibleFunctions []int
	possibleOperators []int
}

func initializeMemoryForModularProgram() []uint64 {
	memory := make([]uint64, len(Constants))
	memory[ONE] = 1
	memory[THREE] = 3

	return memory
}

func newModularProgram(byModule uint64) *programModular {
	return &programModular{
		memory:     initializeMemoryForModularProgram(),
		operations: []Operation{},
		byModule:   byModule,
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
			int(GCD),
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

	program.loadFromExpressionTreeRecursive(tree)

	return program, nil
}

func (program *programModular) loadFromExpressionTreeRecursive(node *Expression) (int, error) {
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

func (program *programModular) GetPointersToFunctionsTypes() []*int {
	result := []*int{}
	l := len(program.operations)
	for i := 0; i < l; i++ {
		if program.operations[i].Operand2Offset == 0 {
			result = append(result, (*int)(&program.operations[i].OperationType))
		}
	}
	return result
}

func (program *programModular) GetPointersToOperatorsTypes() []*int {
	result := []*int{}
	l := len(program.operations)
	for i := 0; i < l; i++ {
		if program.operations[i].Operand2Offset != 0 {
			result = append(result, (*int)(&program.operations[i].OperationType))
		}
	}
	return result
}

func (program *programModular) GetPointersToConstantsOffsets() []*int {
	constantsRange := len(Constants)
	result := []*int{}

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

func (program *programModular) NewFunc(operationType OperationType, operand1Offset int) int {

	program.memory = append(program.memory, 666)
	resultAddr := len(program.memory) - 1

	newOp := Operation{
		Operand1Offset: operand1Offset,
		OperationType:  operationType,
	}

	program.operations = append(program.operations, newOp)

	return resultAddr
}

func (program *programModular) NewOp(operationType OperationType, operand1Offset int, operand2Offset int) int {
	program.memory = append(program.memory, 666)

	newOp := Operation{
		Operand1Offset: operand1Offset,
		Operand2Offset: operand2Offset,
		OperationType:  operationType,
	}
	program.operations = append(program.operations, newOp)
	return len(program.memory) - 1
}

func (program *programModular) toString(operation *Operation) string {
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
func Pow_uint64_mod(x uint64, n uint64, m uint64) uint64 {
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
func Sub_uint64(a uint64, b uint64, m uint64) uint64 {
	return (m + a - b) % m
}

func GCD_uint64(a uint64, b uint64) uint64 {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}

func Pow_uint64(n uint64, m int) uint64 {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func (program *programModular) Execute() uint64 {

	memory := program.memory
	resultsOffset := len(Constants)

	for operationNumber, operation := range program.operations {
		memoryResultOffset := operationNumber + resultsOffset

		switch operation.OperationType {
		case SUM:
			memory[memoryResultOffset] = (memory[operation.Operand1Offset] + memory[operation.Operand2Offset]) % program.byModule
		case SUB:
			memory[memoryResultOffset] = Sub_uint64(memory[operation.Operand1Offset], memory[operation.Operand2Offset], program.byModule)
		case MUL:
			memory[memoryResultOffset] = (memory[operation.Operand1Offset] * memory[operation.Operand2Offset]) % program.byModule
		case POW:
			memory[memoryResultOffset] = Pow_uint64_mod(memory[operation.Operand1Offset], memory[operation.Operand2Offset], program.byModule)
		case GCD:
			memory[memoryResultOffset] = GCD_uint64(memory[operation.Operand1Offset], memory[operation.Operand2Offset])

		default:
			panic(fmt.Sprintf("unknown operationType=%v", operation.OperationType))
		}

		//fmt.Printf("%v %v %v -> %v\n", memory[operation.Operand1Offset], operations[operation.OperationType], memory[operation.Operand2Offset], memory[memoryResultOffset])
	}

	return program.memory[len(program.memory)-1]
}

// Recombine function
// ready(result) is the function which obtain calculation result, if this function returns
func (program *programModular) Recombine(xValues []uint64, maxXOccurrences int, ready func()) {

	constants := program.GetPointersToConstantsOffsets()
	functions := program.GetPointersToFunctionsTypes()
	operations := program.GetPointersToOperatorsTypes()

	for _, x := range xValues {
		program.SetX(x)

		ready_X_Constants_Functions := func() {
			RecombineValues(&operations, &program.possibleOperators, ready)
		}

		ready_X_Constants := func() {
			RecombineValues(&functions, &program.possibleFunctions, ready_X_Constants_Functions)
		}

		readyX := func(remainedConstants *[]*int) {
			RecombineValues(remainedConstants, &program.possibleConstants, ready_X_Constants)
		}

		RecombineRequiredX(&constants, maxXOccurrences, int(X), readyX)
	}

}

func (program *programModular) GetEstimation(maxXOccurrences int) uint64 {
	var sum uint64 = 0

	for i := 1; i <= maxXOccurrences; i++ {
		sum += Combination_uint64(len(program.GetPointersToConstantsOffsets()), i) * // x
			Pow_uint64(uint64(len(program.possibleConstants)), len(program.GetPointersToConstantsOffsets())-i) * // remained constants
			Pow_uint64(uint64(len(program.possibleFunctions)), len(program.GetPointersToFunctionsTypes())) * // functions
			Pow_uint64(uint64(len(program.possibleOperators)), len(program.GetPointersToOperatorsTypes())) // operators
	}

	return sum
}

func Fact_uint64(x int) uint64 {
	result := uint64(1)
	for i := 1; i <= x; i++ {
		result *= uint64(i)
	}
	return result
}

func Combination_uint64(n, k int) uint64 {
	return Fact_uint64(n) / (Fact_uint64(k) * Fact_uint64(n-k))
}
