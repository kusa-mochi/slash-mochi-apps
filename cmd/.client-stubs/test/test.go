package test_service_tester

import (
	"context"
	"net/http"
	"slash_mochi/cmd/.client-stubs/test_kit"
	testv1 "slash_mochi/gen/go/slash_mochi/v1/test"
	"slash_mochi/gen/go/slash_mochi/v1/test/testv1connect"

	"connectrpc.com/connect"
)

type TestServiceTester struct {
	connectClient testv1connect.TestServiceClient
}

func NewTestServiceTester(targetUrl string) *TestServiceTester {
	return &TestServiceTester{
		connectClient: testv1connect.NewTestServiceClient(
			http.DefaultClient,
			targetUrl,
		),
	}
}

func (t *TestServiceTester) Test() []test_kit.TestResult {
	return test_kit.RunTests(
		t.testLoopback_Normal1(),
		t.testLoopback_Normal2(),
	)
}

func (t *TestServiceTester) testLoopback_Normal1() test_kit.TestResult {
	// create new test result object.
	ret := test_kit.NewTestResult()

	// do test.
	res, err := t.connectClient.Loopback(
		context.Background(),
		connect.NewRequest(&testv1.LoopbackRequest{
			Message: "test",
		}),
	)
	if err != nil {
		return ret
	}

	// check result.
	ret.IsSucceeded = res.Msg.Message == "response:test"

	return ret
}

func (t *TestServiceTester) testLoopback_Normal2() test_kit.TestResult {
	ret := test_kit.NewTestResult()

	res, err := t.connectClient.Loopback(
		context.Background(),
		connect.NewRequest(&testv1.LoopbackRequest{
			Message: "hello",
		}),
	)
	if err != nil {
		return ret
	}
	ret.IsSucceeded = res.Msg.Message == "response:hello"

	return ret
}
