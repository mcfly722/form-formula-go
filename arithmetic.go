package formFormula

import "fmt"

type OffsetMEM uint

const (
	UNDEFINED OffsetMEM = 0
	ONE       OffsetMEM = 1
	X         OffsetMEM = 2
	I         OffsetMEM = 3
	PV0       OffsetMEM = 4
	PV1       OffsetMEM = 5
	PVX       OffsetMEM = 6
	PI        OffsetMEM = 7
	E         OffsetMEM = 8
	MINUS_ONE OffsetMEM = 9
	THREE     OffsetMEM = 10
)

type OperationType uint

const (
	NOTHING OperationType = 0
	SUM     OperationType = 1
	SUB     OperationType = 2
	MUL     OperationType = 3
	DIV     OperationType = 4
	FCT     OperationType = 5
	POW     OperationType = 6
	GCD     OperationType = 7
	INVERSE OperationType = 8
	SQRT    OperationType = 9
)

type Operation struct {
	Operand1Offset uint
	Operand2Offset uint
	OperationType  OperationType
}

var Constants = map[OffsetMEM]string{
	UNDEFINED: "undefined",
	X:         "x",
	I:         "i",
	PV0:       "Pv0",
	PV1:       "Pv1",
	PVX:       "PvX",
	ONE:       "1",
	MINUS_ONE: "-1",
	THREE:     "3",
	PI:        "Pi",
	E:         "e",
}

var operations = map[OperationType](func(val1 string, val2 string) string){
	NOTHING: func(val1 string, val2 string) string { return fmt.Sprintf("nothing(%v)", val1) },
	SUM:     func(val1 string, val2 string) string { return fmt.Sprintf("%v+%v", val1, val2) },
	SUB:     func(val1 string, val2 string) string { return fmt.Sprintf("%v-%v", val1, val2) },
	MUL:     func(val1 string, val2 string) string { return fmt.Sprintf("%v*%v", val1, val2) },
	DIV:     func(val1 string, val2 string) string { return fmt.Sprintf("%v/%v", val1, val2) },
	FCT:     func(val1 string, val2 string) string { return fmt.Sprintf("%v!", val1) },
	POW:     func(val1 string, val2 string) string { return fmt.Sprintf("%v^%v", val1, val2) },
	GCD:     func(val1 string, val2 string) string { return fmt.Sprintf("gcd(%v,%v)", val1, val2) },
	INVERSE: func(val1 string, val2 string) string { return fmt.Sprintf("inverse(%v)", val1) },
	SQRT:    func(val1 string, val2 string) string { return fmt.Sprintf("sqrt(%v)", val1) },
}

const defaultErrorUIntValue = ^uint(0)

func Internal_Fact_uint64(x uint) uint64 {
	result := uint64(1)
	for i := uint(1); i <= x; i++ {
		result *= uint64(i)
	}
	return result
}

func Internal_Combination_uint64(n, k uint) uint64 {
	return Internal_Fact_uint64(n) / (Internal_Fact_uint64(k) * Internal_Fact_uint64(n-k))
}

func Internal_Pow_uint64(n uint, m uint) uint64 {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return uint64(n)
	}

	result := uint64(n)
	for i := uint(2); i <= m; i++ {
		result *= uint64(n)
	}
	return result
}

func Internal_GetEstimation(
	maxXOccurrences uint,
	possibleConstants uint,
	possibleFunctions uint,
	possibleOperators uint,
	amountOfPointersToConstantsOffsets uint,
	amountOfPointersToFunctionsTypes uint,
	amountOfPointersToOperatorsTypes uint,
) uint64 {
	var sum uint64 = 0

	for i := uint(1); i <= maxXOccurrences; i++ {
		sum += Internal_Combination_uint64(amountOfPointersToConstantsOffsets, i) * // x
			Internal_Pow_uint64(possibleConstants, amountOfPointersToConstantsOffsets-i) * // remained constants
			Internal_Pow_uint64(possibleFunctions, amountOfPointersToFunctionsTypes) * // functions
			Internal_Pow_uint64(possibleOperators, amountOfPointersToOperatorsTypes) // operators
	}

	return sum
}

/*
func max(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}
*/

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}
