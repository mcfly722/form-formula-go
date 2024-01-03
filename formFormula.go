package formFormula

//#include "vm.h"
import "C"

import (
	"fmt"
)

//export go_vm_debugger
func go_vm_debugger() {
	fmt.Printf("result")
}

func SearchByForm() {

	var memory = [64]C.double{}

	C.Execute(&memory[0], C.closure(C.go_vm_debugger))

}
