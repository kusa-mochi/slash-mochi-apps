package main

import (
	"fmt"
	test_service_tester "slash_mochi/cmd/.client-stubs/test"
	"slash_mochi/cmd/.client-stubs/test_common"
)

func RunTest(runner test_common.Tester) {
	results := runner.Test()
	nResult := len(results)
	for iResult := 0; iResult < nResult; iResult++ {
		if results[iResult].IsSucceeded {
			fmt.Printf("\x1b32mok\x1b[0m|%s\n", results[iResult].TestName)
		} else {
			fmt.Printf("\x1b31mng\x1b[0m|%s\n", results[iResult].TestName)
		}
	}
}

func main() {
	testServiceTester := test_service_tester.NewTestServiceTester()
	RunTest(testServiceTester)
}
