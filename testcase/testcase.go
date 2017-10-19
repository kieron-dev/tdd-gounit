package testcase

import (
	"fmt"
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
	fmt.Println("Num methods", st.NumMethod())
	method, ok := st.MethodByName(tc.method)
	if ok {
		method.Func.Call([]reflect.Value{reflect.ValueOf(tc.testobj)})
	}
}
