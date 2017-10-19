package main

import (
	"testing"

	"github.com/kieron-pivotal/xunit/testcase"
	"github.com/kieron-pivotal/xunit/wasrun"
)

type TestTestCase struct {
	t      *testing.T
	wasRun wasrun.WasRun
}

func (tc *TestTestCase) SetUp() {
	tc.wasRun = wasrun.WasRun{}
}

func (tc *TestTestCase) TestWasRun() {
	test := testcase.New(&tc.wasRun, "TestMethod")
	test.Run()
	if !tc.wasRun.RunFlag {
		tc.t.Error("Method wasn't run")
	}
}

func (tc *TestTestCase) TestSetUp() {
	test := testcase.New(&tc.wasRun, "TestMethod")
	test.Run()
	if !tc.wasRun.SetUpFlag {
		tc.t.Error("SetUp wasn't run")
	}
}

func TestMyTestCase(t *testing.T) {
	testcase.New(&TestTestCase{t: t}, "TestWasRun").Run()
	testcase.New(&TestTestCase{t: t}, "TestSetUp").Run()
}
