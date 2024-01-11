package formFormula

type OffsetMEM int

const (
	UNDEFINED OffsetMEM = 0
	X         OffsetMEM = 1
	I         OffsetMEM = 2
	PV0       OffsetMEM = 3
	PV1       OffsetMEM = 4
	PVX       OffsetMEM = 5
	PI        OffsetMEM = 6
	E         OffsetMEM = 7
	ONE       OffsetMEM = 8
	MINUS_ONE OffsetMEM = 9
	THREE     OffsetMEM = 10
)

type OperationType int

const (
	SUM OperationType = 0
	SUB OperationType = 1
	MUL OperationType = 2
	DIV OperationType = 3
	FCT OperationType = 4
	POW OperationType = 5
	GCD OperationType = 6
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

var operations = map[OperationType]string{
	SUM: "+",
	SUB: "-",
	MUL: "*",
	DIV: "/",
	FCT: "!",
	POW: "^",
	GCD: "gcd",
}
