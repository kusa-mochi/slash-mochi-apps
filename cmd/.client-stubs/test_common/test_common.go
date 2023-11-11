package test_common

type TestResult struct {
	TestName    string
	IsSucceeded bool
}

type Tester interface {
	Test() []TestResult
}
