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
			fmt.Printf("o | %s\n", results[iResult].TestName)
		} else {
			fmt.Printf("x | %s\n", results[iResult].TestName)
		}
	}
}

func main() {
	testServiceTester := test_service_tester.NewTestServiceTester()
	RunTest(testServiceTester)
}
