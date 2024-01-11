package formFormula

import "fmt"

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
}

type programModular struct {
	memory     []uint64
	operations []Operation
	byModule   uint64
}

func NewModularProgram(byModule uint64) ProgramModular {
	memory := make([]uint64, len(Constants))

	memory[ONE] = 1
	memory[THREE] = 3

	return &programModular{
		memory:     memory,
		operations: []Operation{},
		byModule:   byModule,
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
func Pow_uint64(x uint64, n uint64, m uint64) uint64 {
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
			memory[memoryResultOffset] = Pow_uint64(memory[operation.Operand1Offset], memory[operation.Operand2Offset], program.byModule)
		case GCD:
			memory[memoryResultOffset] = GCD_uint64(memory[operation.Operand1Offset], memory[operation.Operand2Offset])

		default:
			panic(fmt.Sprintf("unknown operationType=%v", operation.OperationType))
		}

		//fmt.Printf("%v %v %v -> %v\n", memory[operation.Operand1Offset], operations[operation.OperationType], memory[operation.Operand2Offset], memory[memoryResultOffset])
	}

	return program.memory[len(program.memory)-1]
}
