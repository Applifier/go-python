package python

// #cgo pkg-config: python-3.6
// #include <Python.h>
//
// void pyRelease(PyObject *o)
// {
//	    Py_DECREF(o);
// }
//
// void pyRetain(PyObject *o)
// {
//	    Py_INCREF(o);
// }
import "C"

// PyObject represents a python object
type PyObject struct {
	ptr *C.PyObject
}

// GetAttrString returns an attribute based on a string key
func (obj *PyObject) GetAttrString(key string) (pyObj *PyObject) {
	ptr := C.PyObject_GetAttrString(obj.ptr, C.CString(key))
	if ptr != nil {
		return &PyObject{ptr: ptr}
	}
	return
}

// ToLong returns a PyLong
func (obj *PyObject) ToLong() *PyLong {
	return &PyLong{PyObject: *obj}
}

// Callable tells if the PyObject is callable
func (obj *PyObject) Callable() bool {
	return C.PyCallable_Check(obj.ptr) == 1
}

// Call call a callable Python object callable_object, with arguments given by the tuple args
func (obj *PyObject) Call(args *PyTuple) *PyObject {
	ptr := C.PyObject_CallObject(obj.ptr, args.ptr)
	if ptr != nil {
		return &PyObject{ptr: ptr}
	}

	return nil
}

// Release decreases retain count for PyObject
func (obj *PyObject) Release() {
	if obj.ptr != nil {
		C.pyRetain(obj.ptr)
	}
}

// Retain increases retain count for PyObject
func (obj *PyObject) Retain() {
	if obj.ptr != nil {
		C.pyRelease(obj.ptr)
	}
}
