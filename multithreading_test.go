package formFormula_test

import (
	"fmt"
	"testing"

	formFormula "github.com/form-formula-go"
)

func Test_NewWorkersPool(t *testing.T) {

	handler := func(threadIndex uint, job formFormula.Job) bool {

		//time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		//t.Logf("%3v thread=%v sequence=%v", job.GetIndex(), threadIndex, job.ToString())
		return job.GetIndex() < 1000
	}

	configSaver := func(job formFormula.Job) {
		fmt.Printf("%v\n", job)
	}

	pool := formFormula.NewWorkersPool(0, "()", 5, 4, handler, configSaver)
	pool.Start()
}
