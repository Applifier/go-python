package python

// #include <Python.h>
import "C"

// PyDict represent a python dictionary in go
type PyDict struct {
	PyObject
}

// NewDict creates a new python dictionary
func NewDict() *PyDict {
	ptr := C.PyDict_New()
	if ptr != nil {
		return &PyDict{PyObject{ptr: ptr}}
	}

	return nil
}

// SetItem sets an item to dictionary
func (d *PyDict) SetItem(key string, val *PyObject) {
	C.PyDict_SetItem(
		d.ptr,
		C.PyUnicode_FromString(C.CString(key)),
		val.ptr)
}
