package main

import (
	"fmt"

	"github.com/kieron-pivotal/xunit/testcase"
	"github.com/kieron-pivotal/xunit/wasrun"
)

func main() {
	wasRun := wasrun.WasRun{}
	test := testcase.New(&wasRun, "TestMethod")
	fmt.Println(wasRun.RunFlag)
	test.Run()
	fmt.Println(wasRun.RunFlag)
}
