package main

import (
	"testing"

	"github.com/kieron-pivotal/xunit/testcase"
	"github.com/kieron-pivotal/xunit/wasrun"
)

type TestTestCase struct {
	t *testing.T
}

func (tc *TestTestCase) TestWasRun() {
	wasRun := wasrun.WasRun{}
	test := testcase.New(&wasRun, "TestMethod")
	test.Run()
	if !wasRun.RunFlag {
		tc.t.Error("Method wasn't run")
	}
}

func TestMyTestCase(t *testing.T) {
	testWasRun := testcase.New(&TestTestCase{t}, "TestWasRun")
	testWasRun.Run()
}
