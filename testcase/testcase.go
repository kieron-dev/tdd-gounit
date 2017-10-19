package testcase

import (
	"reflect"
)

type TestCase struct {
	testobj interface{}
	method  func()
}

func New(testobj interface{}, method func()) *TestCase {
	return &TestCase{testobj, method}
}

func (tc *TestCase) Run() {
	st := reflect.TypeOf(tc.testobj)
	args := []reflect.Value{reflect.ValueOf(tc.testobj)}
	if setup, ok := st.MethodByName("SetUp"); ok {
		setup.Func.Call(args)
	}
	tc.method()
}
