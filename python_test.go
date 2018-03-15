package python

import (
	"testing"
)

const fibPython = `
def F(n):
  if n == 0: return 0
  elif n == 1: return 1
  else: return F(n-1)+F(n-2)
`

func fib(n int) uint {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func TestInit(t *testing.T) {
	Initialize()
	res := Run("123")
	if res != 0 {
		t.Error("Error received")
	}
}

func TestRunError(t *testing.T) {
	Initialize()
	res := Run("foo")
	if res != -1 {
		t.Error("Error was not received")
	}
}

func TestGetModule(t *testing.T) {
	Initialize()
	main := AddModule("__main__")
	if main == nil {
		t.Error("Could not get main")
	}
}

func TestGetModuleFromMain(t *testing.T) {
	Initialize()
	main := AddModule("__main__")
	if main == nil {
		t.Error("Could not get main")
	}

	Run("foo = 1")
	foo := main.GetAttrString("foo")
	if foo == nil {
		t.Error("Foo not found")
	}
}

func TestFunctionCall(t *testing.T) {
	Initialize()
	main := AddModule("__main__")
	if main == nil {
		t.Error("Could not get main")
	}

	Run(fibPython)
	fibFunc := main.GetAttrString("F")
	if fibFunc == nil {
		t.Error("fibFunc not found")
	}

	if !fibFunc.Callable() {
		t.Error("Func should be callable")
	}

	tuple := NewTuple(1)
	defer tuple.Release()
	val := NewLongFromInt(int64(10))
	defer val.Release()

	tuple.SetItem(0, val)

	res := fibFunc.Call(tuple)
	if res == nil {
		t.Error("No result received")
	}
	defer res.Release()

	resInt := res.ToLong().AsInt()
	if resInt != int64(fib(10)) {
		t.Error("Wrong value received", resInt)
	}
}

func BenchmarkFunctionCall(b *testing.B) {
	Initialize()
	main := AddModule("__main__")
	if main == nil {
		b.Error("Could not get main")
	}

	Run(fibPython)
	fibFunc := main.GetAttrString("F")
	if fibFunc == nil {
		b.Error("fibFunc not found")
	}

	if !fibFunc.Callable() {
		b.Error("Func should be callable")
	}

	b.ResetTimer()

	tuple := NewTuple(1)
	val := NewLongFromInt(int64(10))
	tuple.SetItem(0, val)

	defer tuple.Release()
	defer val.Release()

	for i := 0; i < b.N; i++ {
		fibFunc.Call(tuple).Release()
	}

}

func BenchmarkGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(10)
	}
}

func BenchmarkPython(b *testing.B) {
	Initialize()
	Run(fibPython)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Run("F(10)")
		//test(fmt.Sprintf("sum *= %d", i))
		//test("\"foo\" + \"bar\" + \"bar\" + \"bar\" + \"bar\" + \"bar\" + \"bar\" + \"bar\" + \"bar\" + \"bar\" + \"bar\" + \"bar\" + \"bar\" + \"bar\"")
	}
}
