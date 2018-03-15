package python

// #include <Python.h>
import "C"

// PyTuple wraps a python tuple
type PyTuple struct {
	PyObject
}

// SetItem sets value to a specific index in the tuple
func (t *PyTuple) SetItem(i int, val *PyObject) {
	C.PyTuple_SetItem(t.ptr, C.long(i), val.ptr)
}

// NewTuple creates a new tuple
func NewTuple(size int) *PyTuple {
	ptr := C.PyTuple_New(C.long(size))
	if ptr != nil {
		return &PyTuple{
			PyObject: PyObject{
				ptr: ptr,
			},
		}
	}

	return nil
}
