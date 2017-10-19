package testcase

import (
	"reflect"
)

type TestCase struct {
	testobj interface{}
	method  string
}

func New(testobj interface{}, method string) *TestCase {
	return &TestCase{testobj, method}
}

func (tc *TestCase) Run() {
	st := reflect.TypeOf(tc.testobj)
	args := []reflect.Value{reflect.ValueOf(tc.testobj)}
	if setup, ok := st.MethodByName("SetUp"); ok {
		setup.Func.Call(args)
	}
	if method, ok := st.MethodByName(tc.method); ok {
		method.Func.Call(args)
	}
}
