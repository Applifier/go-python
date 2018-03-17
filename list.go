package python

// #include <Python.h>
import "C"
import "errors"

// ErrCouldNotInsert could not insert item to list
var ErrCouldNotInsert = errors.New("could not insert item to list")

// PyList represents a python list
type PyList struct {
	PyObject
}

// NewList returns a new python list
func NewList(length int) *PyList {
	ptr := C.PyList_New(C.long(length))
	if ptr != nil {
		return &PyList{PyObject{ptr}}
	}

	return nil
}

// Insert inserts an item to a specified index
func (list *PyList) Insert(index int, obj *PyObject) error {
	if C.PyList_Insert(list.ptr, C.long(index), obj.ptr) == -1 {
		return ErrCouldNotInsert
	}

	return nil
}

// Append appends an item to the list
func (list *PyList) Append(obj *PyObject) error {
	if C.PyList_Append(list.ptr, obj.ptr) == -1 {
		return ErrCouldNotInsert
	}

	return nil
}

// Size returns the size of the list
func (list *PyList) Size() int {
	return int(C.PyList_Size(list.ptr))
}

// GetItem returns an item from the list
func (list *PyList) GetItem(index int) *PyObject {
	ptr := C.PyList_GetItem(list.ptr, C.long(index))
	if ptr != nil {
		return &PyObject{ptr}
	}

	return nil
}

// SetItem set item into given index
func (list *PyList) SetItem(index int, obj *PyObject) error {
	if C.PyList_SetItem(list.ptr, C.long(index), obj.ptr) == -1 {
		return ErrCouldNotInsert
	}

	return nil
}
