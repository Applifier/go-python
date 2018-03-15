package python

// #cgo pkg-config: python-3.6
// #include <Python.h>
import "C"

// PyFloat wraps a python long value
type PyFloat struct {
	PyObject
}

// AsFloat64 convert to float64
func (pl *PyFloat) AsFloat64() float64 {
	return float64(C.PyFloat_AsDouble(pl.ptr))
}

// NewFloatFromDouble creates a new long
func NewFloatFromDouble(val int64) *PyObject {
	ptr := C.PyFloat_FromDouble(C.double(val))
	if ptr != nil {
		return &PyObject{ptr: ptr}
	}

	return nil
}
