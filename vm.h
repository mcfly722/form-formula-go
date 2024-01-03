#define _MEMORY_SIZE 64
#define _OFFSET_ZERO 0
#define _OFFSET_I    1
#define _OFFSET_PV0  2
#define _OFFSET_PV1  3
#define _OFFSET_PVX  4

// sample: https://stackoverflow.com/questions/37157379/passing-function-pointer-to-the-c-code-using-cgo
typedef void (*closure)();
void go_vm_debugger();

void Execute(double *memory,void (*go_vm_debugger)());
