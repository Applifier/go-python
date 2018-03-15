package python

// #cgo pkg-config: python-3.6
// #include <Python.h>
import "C"
import "sync"

var initOnce sync.Once

// Initialize initializes python runtime
func Initialize() {
	initOnce.Do(func() {
		C.Py_Initialize()
	})
}

/*
TODO
func ErrOccured() bool {
	return C.PyErr_Occurred() == 1
}
*/

// Run executes python code
func Run(run string) int {
	return int(C.PyRun_SimpleStringFlags(C.CString(run), nil))
}

// AddModule adds module and returns object assosiated to it
func AddModule(str string) *PyObject {
	ptr := C.PyImport_AddModule(C.CString(str))
	if ptr != nil {
		return &PyObject{ptr: ptr}
	}

	return nil
}
