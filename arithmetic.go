package formFormula

import "fmt"

type OffsetMEM int

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

type OperationType int

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
)

type Operation struct {
	Operand1Offset int
	Operand2Offset int
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
	NOTHING: func(val1 string, val2 string) string { return fmt.Sprintf("%v", val1) },
	SUM:     func(val1 string, val2 string) string { return fmt.Sprintf("%v+%v", val1, val2) },
	SUB:     func(val1 string, val2 string) string { return fmt.Sprintf("%v-%v", val1, val2) },
	MUL:     func(val1 string, val2 string) string { return fmt.Sprintf("%v*%v", val1, val2) },
	DIV:     func(val1 string, val2 string) string { return fmt.Sprintf("%v/%v", val1, val2) },
	FCT:     func(val1 string, val2 string) string { return fmt.Sprintf("%v!", val1) },
	POW:     func(val1 string, val2 string) string { return fmt.Sprintf("%v^%v", val1, val2) },
	GCD:     func(val1 string, val2 string) string { return fmt.Sprintf("gcd(%v,%v)", val1, val2) },
	INVERSE: func(val1 string, val2 string) string { return fmt.Sprintf("inverse(%v)", val1) },
}
