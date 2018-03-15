package python

// #include <Python.h>
import "C"

// PyLong wraps a python long value
type PyLong struct {
	PyObject
}

// AsInt convert to int
func (pl *PyLong) AsInt() int64 {
	return int64(C.PyLong_AsLong(pl.ptr))
}

// AsFloat64 convert to float64
func (pl *PyLong) AsFloat64() float64 {
	return float64(C.PyLong_AsDouble(pl.ptr))
}

// NewLongFromInt creates a new long
func NewLongFromInt(val int64) *PyObject {
	ptr := C.PyLong_FromLong(C.long(val))
	if ptr != nil {
		return &PyObject{ptr: ptr}
	}

	return nil
}
