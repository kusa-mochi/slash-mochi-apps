package test_kit

import (
	"runtime"
	"strings"
)

type TestResult struct {
	TestName    string
	IsSucceeded bool
}

func NewTestResult() TestResult {
	return TestResult{
		TestName:    GetCurrentMethodNameWithLevel(2),
		IsSucceeded: false,
	}
}

type Tester interface {
	Test() []TestResult
}

func RunTests(r ...TestResult) []TestResult {
	return r
}

func GetCurrentMethodNameWithLevel(level int) string {
	pc, _, _, ok := runtime.Caller(level)
	if !ok {
		return ""
	}
	f := runtime.FuncForPC(pc)
	if f == nil {
		return ""
	}
	fullName := f.Name()
	splittedFullName := strings.Split(fullName, ".")
	lenSplittedFullName := len(splittedFullName)
	if lenSplittedFullName == 0 {
		return ""
	}

	return splittedFullName[lenSplittedFullName-1]
}

func GetCurrentMethodName() string {
	return GetCurrentMethodNameWithLevel(1)
}
