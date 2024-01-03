#include <vm.h>

void Execute(double *memory,void (*go_vm_debugger)()) {
	int n;

	for(n=0;n<=_MEMORY_SIZE;n++){
		memory[n]++;
	}

	go_vm_debugger();
}
