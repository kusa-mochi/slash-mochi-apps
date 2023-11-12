package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	omikuji_service_tester "slash_mochi/cmd/.client-stubs/omikuji"
	test_service_tester "slash_mochi/cmd/.client-stubs/test"
	"slash_mochi/cmd/.client-stubs/test_kit"
)

type Target struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
}

type TestConfig struct {
	Target Target `json:"target"`
}

func RunTest(runner test_kit.Tester) {
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
	rawJson, err := os.ReadFile("./test-config.json")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	var testConfig TestConfig
	json.Unmarshal(rawJson, &testConfig)

	targetUrl := fmt.Sprintf("http://%s:%v", testConfig.Target.Ip, testConfig.Target.Port)

	testServiceTester := test_service_tester.NewTestServiceTester(targetUrl)
	omikujiServiceTester := omikuji_service_tester.NewOmikujiServiceTester(targetUrl)

	RunTest(testServiceTester)
	RunTest(omikujiServiceTester)
}
