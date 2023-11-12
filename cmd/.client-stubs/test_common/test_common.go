package test_common

import (
	"runtime"
	"strings"
)

type TestResult struct {
	TestName    string
	IsSucceeded bool
}

type Tester interface {
	Test() []TestResult
}

func GetCurrentMethodName() string {
	pc, _, _, ok := runtime.Caller(1)
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
