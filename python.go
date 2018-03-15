package python

// #cgo pkg-config: python-3.6
// #include <Python.h>
import "C"
import "sync"

var initOnce sync.Once

func newRandomObject() *C.PyObject {
	var dict = C.PyDict_New()
	C.PyDict_SetItem(
		dict,
		C.PyUnicode_FromString(C.CString("key")),
		C.PyUnicode_FromString(C.CString("value")))
	return dict
}

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

func Run(run string) int {
	return int(C.PyRun_SimpleStringFlags(C.CString(run), nil))
}

func AddModule(str string) *PyObject {
	ptr := C.PyImport_AddModule(C.CString(str))
	if ptr != nil {
		return &PyObject{ptr: ptr}
	}

	return nil
}

func EvalGetGlobals() *PyObject {
	ptr := C.PyEval_GetGlobals()
	if ptr != nil {
		return &PyObject{ptr: ptr}
	}

	return nil
}

func EvalGetLocals() *PyObject {
	ptr := C.PyEval_GetLocals()
	if ptr != nil {
		return &PyObject{ptr: ptr}
	}

	return nil
}

func EvalGetBuildins() *PyObject {
	ptr := C.PyEval_GetBuiltins()
	if ptr != nil {
		return &PyObject{ptr: ptr}
	}

	return nil
}
