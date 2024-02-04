package main

import (
	"fmt"
	"math/big"
	"time"

	formFormula "github.com/form-formula-go"
)

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

	// v < max false
	// v > max true
	v bool
}

func main() {
	maxXOccurrences := uint(4)
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

	fmt.Printf("samples: %v\n", samples)

	handler := func(threadIndex uint, job formFormula.Job) bool {

		// form
		program, err := formFormula.NewModularProgramFromBracketsString(p, job.ToString())
		if err != nil {
			panic(err)
		}

		//fmt.Printf("started #%5v: %v estimation=[%v]\n", job.GetIndex(), job.ToString(), program.GetEstimation(maxXOccurrences))

		readyCombination := func() {
			job.IncrementCycle()

			for _, sample := range samples {

				program.SetX(sample.x)

				result := program.Execute()

				// v < max false
				// v > max true

				if result == 0 && sample.v {
					return
				}

				if result != 0 && !sample.v {
					return
				}

				//fmt.Printf("     %30v      x=%8v(%5v) result=%8v\n", program.Disassemble(), sample.x, sample.v, result)
			}

			msg := fmt.Sprintf("solution found: %v", program.Disassemble())

			panic(msg)
		}

		program.Recombine(maxXOccurrences, readyCombination)
		return job.GetIndex() < 300000
	}

	configSaver := func(job formFormula.Job) {
		fmt.Printf("%v\tFORM #%5v: %v  %v\n", time.Now().Format("01-02-2006 15:04:05"), job.GetIndex(), job.ToString(), job.Stat())
	}

	pool := formFormula.NewWorkersPool(0, "()", 100, 700, handler, configSaver)
	pool.Start()
}
