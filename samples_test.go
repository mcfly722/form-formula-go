package formFormula_test

import (
	"fmt"
	"math/big"
	"testing"

	formFormula "github.com/form-formula-go"
)

func Test_isPrime(t *testing.T) {
	for p := int64(1); p < 1000; p++ {
		fmt.Printf("%3v %v\n", p, big.NewInt(p).ProbablyPrime(20))
	}
}

func max(m uint64, n uint64) uint64 {
	if n > m {
		return n
	}
	return m
}

func min(m uint64, n uint64) uint64 {
	if n < m {
		return n
	}
	return m
}

type sample struct {
	x uint64
	v bool
}

func Test_Line(t *testing.T) {
	maxXOccurrences := uint(1)
	n := uint64(31063)
	m := uint64(96179)
	min := min(m, n)
	max := max(m, n)

	p := min * max

	fmt.Printf("%v(min) * %v(max) = %v\n", min, max, p)

	samples := []sample{}

	for i := uint64(2); len(samples) < 200; i = i + 2 {
		nn := max - i
		mm := max + i

		if big.NewInt(int64(nn)).ProbablyPrime(20) {
			samples = append(samples, sample{x: nn, v: false})
		}

		if big.NewInt(int64(mm)).ProbablyPrime(20) {
			samples = append(samples, sample{x: mm, v: true})
		}

	}

	//fmt.Printf("samples: %v\n", samples)

	handler := func(threadIndex uint, job formFormula.Job) bool {

		// form
		fmt.Printf("FORM: %v\n", job.ToString())
		program, err := formFormula.NewModularProgramFromBracketsString(p, job.ToString())
		if err != nil {
			panic(err)
		}

		readyCombination := func() {
			v := program.Disassemble()
			fmt.Printf("%v\n", v)
			program.SetX(samples[0].x)
			result := program.Execute()

			fmt.Printf("     %30v      x=%8v result=%8v\n", program.Disassemble(), samples[0].x, result)
		}

		fmt.Println("recombine start")
		program.Recombine(maxXOccurrences, readyCombination)
		fmt.Println("recombine done")
		return job.GetIndex() < 5
	}

	configSaver := func(job formFormula.Job) {
		//fmt.Printf("%v\n", job)
	}

	pool := formFormula.NewWorkersPool(0, "()", 1, 1, handler, configSaver)
	pool.Start()
}
