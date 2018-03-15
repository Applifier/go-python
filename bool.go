package python

// #include <Python.h>
import "C"

// PyBool represents a bool python bool value
type PyBool struct {
	PyObject
}

// NewPool returns a new python pool object based on a go bool value
func NewPool(val bool) *PyBool {
	longVal := 0
	if val {
		longVal = 1
	}

	ptr := C.PyBool_FromLong(C.long(longVal))

	if ptr != nil {
		return &PyBool{PyObject{ptr: ptr}}
	}

	return nil
}
